package database

import "database/sql"

type Emotion struct {
	Id        int
	Person_id int
	Angry     float32
	Disgust   float32
	Fear      float32
	Happy     float32
	Sad       float32
	Surprise  float32
	Neutral   float32
}

type Client struct {
	db *sql.DB
}

type Person struct {
	Id   int
	Name string
}

type EmotionResponse struct {
	Angry    float32 `json:"angry"`
	Disgust  float32 `json:"disgust"`
	Fear     float32 `json:"fear"`
	Happy    float32 `json:"happy"`
	Sad      float32 `json:"sad"`
	Surprise float32 `json:"surprise"`
	Neutral  float32 `json:"neutral"`
}

type PersonResponse struct {
	Name string `json:"name"`
}

type ServerResponse struct {
	Emotion EmotionResponse `json:"emotions"`
	Person  string          `json:"person"`
}

type ServerArrayResponse struct {
	Persons []ServerResponse `json:"persons"`
}
