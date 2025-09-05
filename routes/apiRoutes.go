package routes

import (
	"enchanted-codex/services"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type apiRouter struct {
	mux           *http.ServeMux
	answerService services.AnswerService
	teamService   services.TeamService
}

func (a apiRouter) createTeamHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, fmt.Sprintf("invalid form data: %s", err), http.StatusBadRequest)
			return
		}
		name := r.Form.Get("team-name")
		if name == "" {
			http.Error(w, "name must not be empty", http.StatusBadRequest)
			return
		}
		members := strings.Split(r.Form.Get("team-members"), ",")
		membersValid := true
		for i, member := range members {
			members[i] = strings.TrimSpace(member)
			if members[i] == "" {
				membersValid = false
			}
		}
		if len(members) < 1 {
			http.Error(w, "must have at least one member", http.StatusBadRequest)
			return
		}
		if !membersValid {
			http.Error(w, "all members must be non-empty strings", http.StatusBadRequest)
			return
		}

		team, err := a.teamService.CreateTeam(services.CreateTeam{
			Name:    name,
			Members: members,
		})
		if err != nil {
			http.Error(w, fmt.Sprintf("error saving team: %s", err), http.StatusInternalServerError)
			return
		}
		id := strconv.Itoa(team.ID)
		teamCookie := http.Cookie{
			Name:     "team",
			Value:    id,
			HttpOnly: true,
			Path:     "/",
			SameSite: http.SameSiteStrictMode,
		}
		http.SetCookie(w, &teamCookie)
		http.Redirect(w, r, "/questions/", http.StatusFound)
	}
}

func (a apiRouter) checkAnswerHandler(part services.Part) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		team, err := getTeamFromContext(true, r.Context())
		if err != nil {
			http.Error(w, fmt.Sprintf("no team on context: %s", err), http.StatusInternalServerError)
			return
		}
		questionID, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		err = r.ParseForm()
		if err != nil {
			http.Error(w, fmt.Sprintf("invalid form data: %s", err), http.StatusBadRequest)
			return
		}
		answerVal := r.Form.Get("answer")
		if answerVal == "" {
			http.Error(w, "answer must not be empty", http.StatusBadRequest)
			return
		}
		reader := csv.NewReader(strings.NewReader(answerVal))
		answerStr := new(strings.Builder)
		for {
			record, err := reader.Read()
			if err != nil {
				if errors.Is(err, io.EOF) {
					break
				}
				http.Error(w, fmt.Sprintf("error reading answer csv: %s", err.Error()), http.StatusBadRequest)
				return
			}
			if answerStr.Len() != 0 {
				answerStr.WriteString(",")
			}
			newStr := new(strings.Builder)
			for i, val := range record {
				newStr.WriteString(strings.TrimSpace(val))
				if i < len(record)-1 {
					newStr.WriteString(",")
				}
			}
			answerStr.WriteString(newStr.String())
		}
		_, err = a.answerService.CheckAnswer(services.AnswerData{
			TeamID:     team.ID,
			QuestionID: questionID,
			Part:       part,
		}, answerStr.String())
		if err != nil {
			http.Error(w, fmt.Sprintf("answer error: %s", err.Error()), http.StatusBadRequest)
			return
		}
		http.Redirect(w, r, fmt.Sprintf("/questions/%d/", questionID), http.StatusFound)
	}
}

func (a apiRouter) Use() {
	a.mux.HandleFunc("POST /create-team/", a.createTeamHandler())
	a.mux.HandleFunc("POST /questions/{id}/part1", checkTeamCookie(a.teamService, true, a.checkAnswerHandler(services.Part1)))
	a.mux.HandleFunc("POST /questions/{id}/part2", checkTeamCookie(a.teamService, true, a.checkAnswerHandler(services.Part2)))
}

func NewAPIRoutes(mux *http.ServeMux, answerService services.AnswerService, teamService services.TeamService) Router {
	return apiRouter{
		mux:           mux,
		answerService: answerService,
		teamService:   teamService,
	}
}
