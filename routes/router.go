package routes

import (
	"context"
	"enchanted-codex/services"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

type ContextKey = int

const (
	teamKey ContextKey = iota
)

type Router interface {
	Use()
}

func getTeamFromContext(required bool, ctx context.Context) (*services.Team, error) {
	ctxVal := ctx.Value(teamKey)
	if ctxVal == nil && !required {
		return nil, nil
	}
	val, ok := ctxVal.(services.Team)
	if !ok {
		return nil, errors.New("invalid or missing team name on context")
	}
	return &val, nil
}

func checkTeamCookie(teamService services.TeamService, required bool, handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("team")
		if err != nil && !errors.Is(err, http.ErrNoCookie) {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if cookie == nil {
			if required {
				http.Error(w, "unauthorized", http.StatusUnauthorized)
			} else {
				handler(w, r)
			}
			return
		}
		id, err := strconv.Atoi(cookie.Value)
		if err != nil {
			http.Error(w, fmt.Sprintf("invalid team cookie: %s", err), http.StatusUnauthorized)
			return
		}
		team, err := teamService.GetTeamByID(id)
		if err != nil {
			http.Error(w, fmt.Sprintf("error fetching team: %s", err), http.StatusUnauthorized)
			return
		}
		r = r.WithContext(context.WithValue(r.Context(), teamKey, *team))
		handler(w, r)
	}
}
