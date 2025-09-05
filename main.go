package main

import (
	"enchanted-codex/database"
	"enchanted-codex/routes"
	"enchanted-codex/services"
	"log"
	"net/http"
)

func LogMiddleware(mux *http.ServeMux) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s: %s", r.Method, r.URL.Path)
		mux.ServeHTTP(w, r)
	}
}

func main() {
	db, err := database.GetDB()
	if err != nil {
		panic(err)
	}
	teamService, questionService, answerService := services.NewTeamServicxe(db), services.NewQuestionService(db), services.NewAnswerService(db)
	mux := new(http.ServeMux)
	viewRouter := routes.NewViewRouter(mux, teamService, questionService, answerService)
	viewRouter.Use()
	apiRouter := routes.NewAPIRoutes(mux, answerService, teamService)
	apiRouter.Use()
	log.Println("Server running on port 4000")
	err = http.ListenAndServe(":4000", LogMiddleware(mux))
	if err != nil {
		panic(err)
	}
}
