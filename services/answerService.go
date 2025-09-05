package services

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type Part = int

const (
	Part1 Part = iota
	Part2
)

type Answer struct {
	ID         int
	TeamID     int
	QuestionID int
	Part       Part
	CreatedAt  time.Time
}

type AnswerData struct {
	TeamID     int
	QuestionID int
	Part       Part
}

type AnswerService interface {
	CheckAnswer(input AnswerData, value string) (*Answer, error)
	HasAnswered(input AnswerData) (*Answer, error)
}

type answerService struct {
	db *sql.DB
}

// HasAnswered implements AnswerService.
func (a answerService) HasAnswered(input AnswerData) (*Answer, error) {
	result := a.db.QueryRow("SELECT id, created_at FROM answers WHERE team_id = ? AND question_id = ? AND part = ?;", input.TeamID, input.QuestionID, input.Part)
	answer := Answer{
		TeamID:     input.TeamID,
		QuestionID: input.QuestionID,
		Part:       input.Part,
	}
	err := result.Scan(&answer.ID, &answer.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &answer, nil
}

// CheckAnswer implements AnswerService.
func (a answerService) CheckAnswer(input AnswerData, value string) (*Answer, error) {
	result := a.db.QueryRow("SELECT id, part_1_answer, part_2_answer FROM questions WHERE id = ?", input.QuestionID)
	var question Question
	err := result.Scan(&question.ID, &question.Part1Answer, &question.Part2Answer)
	if err != nil {
		return nil, err
	}
	fmt.Println(input.Part, value, question.Part1Answer, question.Part2Answer)
	switch input.Part {
	case Part1:
		if value != question.Part1Answer {
			return nil, errors.New("incorrect answer")
		}
	case Part2:
		if value != question.Part2Answer {
			return nil, errors.New("incorrect answer")
		}
	}
	answerResult := a.db.QueryRow("INSERT INTO answers (team_id, question_id, part) VALUES (?, ?, ?) RETURNING id, team_id, question_id, part, created_at", input.TeamID, input.QuestionID, input.Part)
	var answer Answer
	err = answerResult.Scan(&answer.ID, &answer.TeamID, &answer.QuestionID, &answer.Part, &answer.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &answer, nil
}

func NewAnswerService(db *sql.DB) AnswerService {
	return answerService{db}
}
