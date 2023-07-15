package models

import "time"

type Professor struct {
	Id         string    `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Department string    `json:"department"`
	Degree     string    `json:"degree"`
	Subjects   []string  `json:"subjects"`
	Messanger  []Message `json:"messanger"`
}

type Message struct {
	Id       string    `json:"id"`
	SenderId string    `json:"sender_id"`
	Data     string    `json:"data"`
	Time     time.Time `json:"time"`
}
