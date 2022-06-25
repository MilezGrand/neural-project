package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/m4yb3/neural-project-server/internal/database"
	"github.com/m4yb3/neural-project-server/internal/video"

	vidio "github.com/AlexEidt/Vidio"
	"github.com/disintegration/imaging"
)

var client *database.Client

func NewServer(_client *database.Client) error {
	client = _client

	r := mux.NewRouter()

	r.HandleFunc("/emotion/recognize", recognizeEmotion).Methods("POST")
	r.HandleFunc("/database/persons", GetAllPersonsHandler).Methods("GET")
	r.HandleFunc("/database/persons/get", GetPersonHandler).Methods("GET")
	r.HandleFunc("/database/persons/add", AddPersonHandler).Methods("GET")
	r.HandleFunc("/database/emotions/get", GetEmotionHandler).Methods("GET")
	r.HandleFunc("/database/emotions/add", AddEmotionHandler).Methods("POST")

	http_server := &http.Server{
		Handler:      r,
		Addr:         ":49812",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return http_server.ListenAndServe()
}

func GetPersonHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	person, err := client.GetPerson(name)
	if err != nil {
		fmt.Fprintf(w, "err: %v", err)
		return
	}

	bytes, err := json.Marshal(person)
	if err != nil {
		fmt.Fprintf(w, "err: %v", err)
		return
	}

	fmt.Fprintf(w, "%s\n", bytes)
}

func AddPersonHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	err := client.AddPerson(name)
	if err != nil {
		fmt.Fprintf(w, "err: %v", err)
		return
	}

	fmt.Fprintf(w, "True\n")
}

func GetEmotionHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		fmt.Fprintf(w, "err: %v", err)
		return
	}

	var emotion database.Emotion

	emotions, err := client.GetPersonEmotions(id)
	if err != nil {
		fmt.Fprintf(w, "err: %v", err)
		return
	}

	if len(emotions) == 0 {
		fmt.Fprintln(w, "Not found")
		return
	}

	for _, temp_emotion := range emotions {
		emotion.Id = temp_emotion.Id
		emotion.Person_id = temp_emotion.Person_id
		emotion.Angry += temp_emotion.Angry
		emotion.Disgust += temp_emotion.Disgust
		emotion.Fear += temp_emotion.Fear
		emotion.Sad += temp_emotion.Sad
		emotion.Neutral += temp_emotion.Neutral
		emotion.Happy += temp_emotion.Happy
		emotion.Surprise += temp_emotion.Surprise
	}

	emotion.Angry /= float32(len(emotions))
	emotion.Disgust /= float32(len(emotions))
	emotion.Fear /= float32(len(emotions))
	emotion.Sad /= float32(len(emotions))
	emotion.Neutral /= float32(len(emotions))
	emotion.Happy /= float32(len(emotions))
	emotion.Surprise /= float32(len(emotions))

	bytes, err := json.Marshal(emotion)
	if err != nil {
		fmt.Fprintf(w, "err: %v", err)
		return
	}

	fmt.Fprintf(w, "%s\n", bytes)
}

func GetAllPersonsHandler(w http.ResponseWriter, r *http.Request) {
	persons, err := client.GetPersons()
	if err != nil {
		fmt.Fprintf(w, "err: %v", err)
		return
	}

	bytes, err := json.Marshal(persons)
	if err != nil {
		fmt.Fprintf(w, "err: %v", err)
		return
	}

	fmt.Fprintf(w, "%s\n", bytes)
}

func AddEmotionHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		fmt.Fprintf(w, "err: %v", err)
		return
	}

	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		fmt.Fprintf(w, "err: %v", err)
		return
	}

	emotion := r.FormValue("emotion")
	var parsedEmotion database.EmotionResponse

	err = json.Unmarshal([]byte(emotion), &parsedEmotion)
	if err != nil {
		fmt.Fprintf(w, "err: %v", err)
		return
	}

	err = client.AddPersonEmotion(id, parsedEmotion)
	if err != nil {
		fmt.Fprintf(w, "err: %v", err)
		return
	}

	fmt.Fprintf(w, "True\n")
}

func recognizeEmotion(w http.ResponseWriter, r *http.Request) {
	file, err := GetFileFromForm(r)
	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, "err: %v", err)
		return
	}

	defer file.Close()

	filename := file.Name()

	clip, err := vidio.NewVideo(filename)
	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, "err: %v", err)
		return
	}

	fps := int(math.Round(clip.FPS()))
	frames := clip.Frames()
	clip.Close()

	j := 1

	var clips []*Clip

	for i := 1; i < frames; i += fps * 2 {
		reader := video.ExampleReadFrameAsJpeg(filename, i)
		img, err := imaging.Decode(reader)
		if err != nil {
			log.Fatal(err)
		}

		err = imaging.Save(img, fmt.Sprintf("./images/out%d.jpeg", j))
		if err != nil {
			log.Fatal(err)
		}

		content, err := SendPostRequest("http://localhost:51267/recognition/emotion", fmt.Sprintf("./images/out%d.jpeg", j))
		if err != nil {
			fmt.Println(err)
			fmt.Fprintf(w, "err: %v", err)
			return
		}

		var serverResponse database.ServerResponse
		err = json.Unmarshal(content, &serverResponse)
		if err != nil {
			fmt.Println(err)
			fmt.Fprintf(w, "err: %v", err)
			return
		}

		personDB, err := client.GetPerson(serverResponse.Person)
		if err != nil {
			fmt.Println(err)
			fmt.Fprintf(w, "err: %v", err)
			return
		}

		if personDB.Name == "" {
			return
		}

		err = client.AddPersonEmotion(personDB.Id, serverResponse.Emotion)
		if err != nil {
			fmt.Println(err)
			fmt.Fprintf(w, "err: %v", err)
			return
		}

		os.Remove(fmt.Sprintf("./images/out%d.jpeg", j))
		j++

		clip := &Clip{Second: i / fps, Emotions: serverResponse}
		clips = append(clips, clip)
	}

	content, err := json.Marshal(clips)
	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, "err: %v", err)
		return
	}

	fmt.Fprintln(w, string(content))
}

func GetFileFromForm(r *http.Request) (*os.File, error) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		return nil, err
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		return nil, err
	}

	tempFile, err := os.CreateTemp("", "image-*.mp4")
	if err != nil {
		return nil, err
	}

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	tempFile.Write(fileBytes)
	return tempFile, nil
}

func SendPostRequest(urlPath string, filename string) ([]byte, error) {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, err := os.Open(filename)

	defer file.Close()

	part1, err := writer.CreateFormFile("file", filename)
	_, err = io.Copy(part1, file)
	if err != nil {
		return nil, err
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", urlPath, payload)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// 2022/06/25 17:01:29 http: panic serving [::1]:64281: runtime error: invalid memory address or nil pointer dereference
// goroutine 18 [running]:
// net/http.(*conn).serve.func1()
//         C:/Program Files/Go/src/net/http/server.go:1825 +0xbf
// panic({0x7f93c0, 0xcae0d0})
//         C:/Program Files/Go/src/runtime/panic.go:844 +0x258
// os.(*Process).signal(0xc0002fc000?, {0x9f6650?, 0xc77ce0?})
//         C:/Program Files/Go/src/os/exec_windows.go:49 +0x4d
// os.(*Process).Signal(...)
//         C:/Program Files/Go/src/os/exec.go:138
// os.(*Process).kill(...)
//         C:/Program Files/Go/src/os/exec_posix.go:67
// os.(*Process).Kill(...)
//         C:/Program Files/Go/src/os/exec.go:123
// github.com/AlexEidt/Vidio.checkExists({0x891e20, 0x6})
//         C:/Users/norms/go/pkg/mod/github.com/!alex!eidt/!vidio@v1.1.1/utils.go:31 +0x89
// github.com/AlexEidt/Vidio.NewVideo({0xc00001e200, 0x36})
//         C:/Users/norms/go/pkg/mod/github.com/!alex!eidt/!vidio@v1.1.1/video.go:93 +0x45
// github.com/m4yb3/neural-project-server/internal/server.recognizeEmotion({0x9f6d30?, 0xc0002900e0}, 0xc000292300?)
//         C:/Users/norms/Desktop/project/neural-project/golang-main-server/internal/server/server.go:182 +0x197
// net/http.HandlerFunc.ServeHTTP(0xc000298000?, {0x9f6d30?, 0xc0002900e0?}, 0x55feb4?)
//         C:/Program Files/Go/src/net/http/server.go:2084 +0x2f
// github.com/gorilla/mux.(*Router).ServeHTTP(0xc0002820c0, {0x9f6d30, 0xc0002900e0}, 0xc00030e000)
//         C:/Users/norms/go/pkg/mod/github.com/gorilla/mux@v1.8.0/mux.go:210 +0x1cf
// net/http.serverHandler.ServeHTTP({0x9f5a98?}, {0x9f6d30, 0xc0002900e0}, 0xc00030e000)
//         C:/Program Files/Go/src/net/http/server.go:2916 +0x43b
// net/http.(*conn).serve(0xc0000aa000, {0x9f7078, 0xc000292120})
//         C:/Program Files/Go/src/net/http/server.go:1966 +0x5d7
// created by net/http.(*Server).Serve
//         C:/Program Files/Go/src/net/http/server.go:3071 +0x4db
// exit status 0xc000013a

// norms /c/Users/norms/Deskt
