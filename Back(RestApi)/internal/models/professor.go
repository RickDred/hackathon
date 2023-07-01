package models

type Professor struct {
	Id         string   `json:"id"`
	Name       string   `json:"name"`
	Email      string   `json:"email"`
	Department string   `json:"department"`
	Degree     string   `json:"degree"`
	Subjects   []string `json:"subjects"`
}
