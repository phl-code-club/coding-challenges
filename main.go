package main

import (
	"enchanted-codex/database"
	"enchanted-codex/services"
)

func main() {
	db, err := database.GetDB()
	if err != nil {
		panic(err)
	}
	teamService, questionService, answerService := services.NewTeamServicxe(db), services.NewQuestionService(db), services.NewAnswerService(db)
	// Throw these out to fix allow for compilation
	_, _, _ = teamService, questionService, answerService
}
