package services

import "database/sql"

type Question struct {
	ID          int
	Name        string
	Part1Desc   string
	Part1Answer string
	Part2Desc   string
	Part2Answer string
}

type QuestionService interface {
	ListQuestions() ([]Question, error)
	GetQuestionByID(id int) (Question, error)
}

type questionService struct {
	db *sql.DB
}

// GetQuestionByID implements QuestionService.
func (q questionService) GetQuestionByID(id int) (Question, error) {
	panic("unimplemented")
}

// ListQuestions implements QuestionService.
func (q questionService) ListQuestions() ([]Question, error) {
	panic("unimplemented")
}

func NewQuestionService(db *sql.DB) QuestionService {
	return questionService{db}
}
