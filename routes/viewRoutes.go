package routes

import (
	"context"
	"errors"
	"html/template"
	"net/http"
)

type ContextKey = int

const (
	teamName ContextKey = iota
)

type ViewRouter interface {
	Use()
}

type viewRouter struct {
	mux *http.ServeMux
}

func (v viewRouter) getTeamFromContext(ctx context.Context) (string, error) {
	val := ctx.Value(teamName)
	valStr, ok := val.(string)
	if !ok {
		return "", errors.New("invalid or missing team name on context")
	}
	return valStr, nil
}

func (v viewRouter) checkTeamCookie(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("team")
		if err != nil && !errors.Is(err, http.ErrNoCookie) {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if cookie == nil {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		r = r.WithContext(context.WithValue(r.Context(), teamName, cookie.Value))
		handler(w, r)
	}
}

func (v viewRouter) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("./html/layout.html", "./html/index.html", "./html/navbar.html"))
		tmpl.Execute(w, nil)
	}
}

func (v viewRouter) handleCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("./html/layout.html", "./html/create-team.html", "./html/navbar.html"))
		tmpl.Execute(w, nil)
	}
}

func (v viewRouter) handleQuestionList() http.HandlerFunc {
	return v.checkTeamCookie(
		func(w http.ResponseWriter, r *http.Request) {
			tmpl := template.Must(template.ParseFiles("./html/layout.html", "./html/question-list.html", "./html/navbar.html"))
			tmpl.Execute(w, nil)
		},
	)
}

func (v viewRouter) handleQuestion() http.HandlerFunc {
	return v.checkTeamCookie(
		func(w http.ResponseWriter, r *http.Request) {
			tmpl := template.Must(template.ParseFiles("./html/layout.html", "./html/question.html", "./html/navbar.html"))
			tmpl.Execute(w, nil)
		},
	)
}

// Use implements ViewRouter.
func (v viewRouter) Use() {
	mux := v.mux
	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("GET /static/", http.StripPrefix("/static/", fs))
	mux.HandleFunc("GET /", v.handleIndex())
	mux.HandleFunc("GET /create/", v.handleCreate())
	mux.HandleFunc("GET /questions/", v.handleQuestionList())
	mux.HandleFunc("GET /questions/{slug}/", v.handleQuestion())
}

func NewViewRouter(mux *http.ServeMux) ViewRouter {
	return viewRouter{mux}
}
