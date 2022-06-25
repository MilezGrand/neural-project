package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func NewDatabase(user string, password string, ip string, database string) (*Client, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, ip, database))
	return &Client{db}, err
}

func (client *Client) AddPerson(name string) error {
	_, err := client.db.Exec("INSERT INTO persons (name) VALUES (?)", name)
	return err
}

func (client *Client) GetPerson(name string) (*Person, error) {
	user := &Person{}
	err := client.db.QueryRow("SELECT * FROM persons WHERE name = ?", name).Scan(&user.Id, &user.Name)
	return user, err
}

func (client *Client) GetPersons() ([]*Person, error) {
	var users []*Person
	rows, err := client.db.Query("SELECT * FROM persons")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		user := &Person{}
		if err := rows.Scan(&user.Id, &user.Name); err != nil {
			return users, err
		}
		users = append(users, user)
	}

	return users, err
}

func (client *Client) AddPersonEmotion(personId int, emotion EmotionResponse) error {
	_, err := client.db.Exec("INSERT INTO emotions (person_id, angry, disgust, fear, happy, sad, surprise, neutral) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", personId, emotion.Angry, emotion.Disgust, emotion.Fear, emotion.Happy, emotion.Sad, emotion.Surprise, emotion.Neutral)
	return err
}

func (client *Client) GetPersonEmotions(personId int) ([]*Emotion, error) {
	var emotions []*Emotion
	rows, err := client.db.Query("SELECT * FROM emotions WHERE person_id = ?", personId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		emotion := &Emotion{}
		if err := rows.Scan(&emotion.Id, &emotion.Person_id, &emotion.Angry, &emotion.Disgust, &emotion.Fear, &emotion.Happy, &emotion.Sad, &emotion.Surprise, &emotion.Neutral); err != nil {
			return emotions, err
		}
		emotions = append(emotions, emotion)
	}

	return emotions, err
}
