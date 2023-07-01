package models

type Review struct {
	Id             string  `json:"id"`
	StudentId      int64   `json:"student_id"`
	ProfessorId    int64   `json:"professor_id"`
	Communication  []int64 `json:"communication"`
	Materials      []int64 `json:"materials"`
	TimeManagement []int64 `json:"time_management"`
	Grading        []int64 `json:"grading"`
	Overall        []int64 `json:"overall"`
	Feedback       string  `json:"feedback"`
}
