package services

import (
	"database/sql"
	"html/template"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
)

type Question struct {
	ID               int
	Name             string
	Intro            string
	Input            string
	Part1Desc        string
	Part1Answer      string
	Part2Desc        string
	Part2Answer      string
	HasAnsweredPart1 bool
	HasAnsweredPart2 bool
}

func (q Question) LockBoxCode() int {
	switch q.ID {
	case 1:
		return 1
	case 2:
		return 3
	case 3:
		return 3
	case 4:
		return 7
	default:
		return 0
	}
}

func (q Question) ParseMarkdown(input string) template.HTML {
	p := parser.NewWithExtensions(parser.CommonExtensions | parser.FencedCode)
	return template.HTML(markdown.ToHTML([]byte(input), p, nil))
}

type QuestionService interface {
	ListQuestions() ([]*Question, error)
	GetQuestionByID(id int) (*Question, error)
}

type questionService struct {
	db *sql.DB
}

// GetQuestionByID implements QuestionService.
func (q questionService) GetQuestionByID(id int) (*Question, error) {
	result := q.db.QueryRow("SELECT id, name, intro, input, part_1_description, part_1_answer, part_2_description, part_2_answer FROM questions WHERE id = ?", id)
	var question Question
	err := result.Scan(&question.ID, &question.Name, &question.Intro, &question.Input, &question.Part1Desc, &question.Part1Answer, &question.Part2Desc, &question.Part2Answer)
	if err != nil {
		return nil, err
	}
	return &question, nil
}

// ListQuestions implements QuestionService.
func (q questionService) ListQuestions() ([]*Question, error) {
	result, err := q.db.Query("SELECT id, name FROM questions;")
	if err != nil {
		return nil, err
	}
	defer result.Close()
	var questions []*Question

	for result.Next() {
		var question Question
		if err := result.Scan(&question.ID, &question.Name); err != nil {
			return nil, err
		}
		questions = append(questions, &question)
	}

	return questions, nil
}

func NewQuestionService(db *sql.DB) QuestionService {
	return questionService{db}
}
