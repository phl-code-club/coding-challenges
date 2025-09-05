package routes

import (
	"enchanted-codex/services"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type PageData struct {
	Team      *services.Team
	Questions []*services.Question
	Question  *services.Question
}

type viewRouter struct {
	mux             *http.ServeMux
	teamService     services.TeamService
	questionService services.QuestionService
	answerService   services.AnswerService
}

func (v viewRouter) parseTemplateWithLayout(path string) *template.Template {
	return template.Must(template.ParseFiles("./html/layout.html", "./html/navbar.html", path))
}

func (v viewRouter) handleIndex() http.HandlerFunc {
	return checkTeamCookie(
		v.teamService,
		false,
		func(w http.ResponseWriter, r *http.Request) {
			team, err := getTeamFromContext(false, r.Context())
			if err != nil {
				http.Error(w, fmt.Sprintf("no team on context: %s", err), http.StatusInternalServerError)
				return
			}
			tmpl := v.parseTemplateWithLayout("./html/index.html")
			err = tmpl.Execute(w, PageData{
				Team: team,
			})
			if err != nil {
				http.Error(w, fmt.Sprintf("unable to render template: %s", err), http.StatusInternalServerError)
			}
		})
}

func (v viewRouter) handleCreate() http.HandlerFunc {
	return checkTeamCookie(
		v.teamService,
		false, func(w http.ResponseWriter, r *http.Request) {
			team, err := getTeamFromContext(false, r.Context())
			if err != nil {
				http.Error(w, fmt.Sprintf("no team on context: %s", err), http.StatusInternalServerError)
				return
			}
			if team != nil {
				http.Redirect(w, r, "/questions", http.StatusFound)
				return
			}
			tmpl := v.parseTemplateWithLayout("./html/create-team.html")
			err = tmpl.Execute(w, nil)
			if err != nil {
				http.Error(w, fmt.Sprintf("unable to render template: %s", err), http.StatusInternalServerError)
			}
		})
}

func (v viewRouter) handleQuestionList() http.HandlerFunc {
	return checkTeamCookie(
		v.teamService,
		true,
		func(w http.ResponseWriter, r *http.Request) {
			team, err := getTeamFromContext(true, r.Context())
			if err != nil {
				http.Error(w, fmt.Sprintf("no team on context: %s", err), http.StatusInternalServerError)
				return
			}
			questions, err := v.questionService.ListQuestions()
			if err != nil {
				http.Error(w, fmt.Sprintf("unable to load questions: %s", err), http.StatusInternalServerError)
				return
			}
			for _, question := range questions {
				answer1, err := v.answerService.HasAnswered(services.AnswerData{
					TeamID:     team.ID,
					QuestionID: question.ID,
					Part:       services.Part1,
				})
				if err != nil {
					http.Error(w, fmt.Sprintf("unable to load answer: %s", err), http.StatusInternalServerError)
					return
				}
				answer2, err := v.answerService.HasAnswered(services.AnswerData{
					TeamID:     team.ID,
					QuestionID: question.ID,
					Part:       services.Part2,
				})
				if err != nil {
					http.Error(w, fmt.Sprintf("unable to load answer: %s", err), http.StatusInternalServerError)
					return
				}
				question.HasAnsweredPart1 = answer1 != nil
				question.HasAnsweredPart2 = answer2 != nil
			}
			tmpl := v.parseTemplateWithLayout("./html/question-list.html")
			err = tmpl.Execute(w, PageData{
				Team:      team,
				Questions: questions,
			})
			if err != nil {
				http.Error(w, fmt.Sprintf("unable to render template: %s", err), http.StatusInternalServerError)
			}
		},
	)
}

func (v viewRouter) handleQuestion() http.HandlerFunc {
	return checkTeamCookie(
		v.teamService,
		true,
		func(w http.ResponseWriter, r *http.Request) {
			team, err := getTeamFromContext(true, r.Context())
			if err != nil {
				http.Error(w, fmt.Sprintf("no team on context: %s", err), http.StatusInternalServerError)
				return
			}
			id, err := strconv.Atoi(r.PathValue("id"))
			if err != nil {
				http.Error(w, "not found", http.StatusNotFound)
				return
			}
			question, err := v.questionService.GetQuestionByID(id)
			if err != nil {
				http.Error(w, fmt.Sprintf("unable to load question: %s", err), http.StatusInternalServerError)
				return
			}
			answer1, err := v.answerService.HasAnswered(services.AnswerData{
				TeamID:     team.ID,
				QuestionID: id,
				Part:       services.Part1,
			})
			if err != nil {
				http.Error(w, fmt.Sprintf("unable to load answer: %s", err), http.StatusInternalServerError)
				return
			}
			answer2, err := v.answerService.HasAnswered(services.AnswerData{
				TeamID:     team.ID,
				QuestionID: id,
				Part:       services.Part2,
			})
			if err != nil {
				http.Error(w, fmt.Sprintf("unable to load answer: %s", err), http.StatusInternalServerError)
				return
			}
			question.HasAnsweredPart1 = answer1 != nil
			question.HasAnsweredPart2 = answer2 != nil
			tmpl := v.parseTemplateWithLayout("./html/question.html")
			err = tmpl.Execute(w, PageData{
				Team:     team,
				Question: question,
			})
			if err != nil {
				http.Error(w, fmt.Sprintf("unable to render template: %s", err), http.StatusInternalServerError)
			}
		},
	)
}

func (v viewRouter) handleQuestionInput() http.HandlerFunc {
	return checkTeamCookie(
		v.teamService,
		true,
		func(w http.ResponseWriter, r *http.Request) {
			id, err := strconv.Atoi(r.PathValue("id"))
			if err != nil {
				http.Error(w, "not found", http.StatusNotFound)
				return
			}
			question, err := v.questionService.GetQuestionByID(id)
			if err != nil {
				http.Error(w, fmt.Sprintf("unable to load question: %s", err), http.StatusInternalServerError)
				return
			}
			_, err = w.Write([]byte(question.Input))
			w.Header().Add("content-type", "text/plain")
			if err != nil {
				http.Error(w, fmt.Sprintf("unable to get input: %s", err), http.StatusInternalServerError)
			}
		},
	)
}

// Use implements ViewRouter.
func (v viewRouter) Use() {
	mux := v.mux
	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("GET /static/", http.StripPrefix("/static/", fs))
	mux.HandleFunc("GET /", v.handleIndex())
	mux.HandleFunc("GET /create-team/", v.handleCreate())
	mux.HandleFunc("GET /questions/", v.handleQuestionList())
	mux.HandleFunc("GET /questions/{id}/", v.handleQuestion())
	mux.HandleFunc("GET /questions/{id}/input", v.handleQuestionInput())
}

func NewViewRouter(mux *http.ServeMux, teamService services.TeamService, questionService services.QuestionService, answerService services.AnswerService) Router {
	return viewRouter{
		mux:             mux,
		teamService:     teamService,
		questionService: questionService,
		answerService:   answerService,
	}
}
