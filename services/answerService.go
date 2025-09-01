package services

import (
	"database/sql"
	"time"
)

type Answer struct {
	ID         int
	TeamID     int
	QuestionID int
	Value      string
	CreatedAt  time.Time
}

type CheckAnswer struct {
	TeamID     int
	QuestionID int
	Value      string
}

type AnswerService interface {
	CheckAnswer(input CheckAnswer) (Answer, error)
}

type answerService struct {
	db *sql.DB
}

// CheckAnswer implements AnswerService.
func (a answerService) CheckAnswer(input CheckAnswer) (Answer, error) {
	panic("unimplemented")
}

func NewAnswerService(db *sql.DB) AnswerService {
	return answerService{db}
}
