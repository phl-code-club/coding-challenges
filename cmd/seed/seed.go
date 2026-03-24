package main

import (
	"database/sql"
	"enchanted-codex/cmd/seed/questions"
	"enchanted-codex/database"
	"errors"
	"fmt"
	"io/fs"
	"os"
)

func main() {
	questions := []questions.Question{
		questions.GenerateQuestion1(),
		questions.GenerateQuestion2(),
		questions.GenerateQuestion3(),
		questions.GenerateQuestion4(),
	}

	err := os.Remove("./test.db")
	if err != nil && !errors.Is(err, fs.ErrNotExist) {
		panic(err)
	}
	db, err := database.GetDB()
	if err != nil {
		panic(err)
	}

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	err = insertQuestions(tx, questions)
	if err != nil {
		handleError(tx, err, "no questions")
		return
	}

	err = tx.Commit()
	if err != nil {
		panic(err)
	}
	fmt.Println("Database seeded successfully!")
}

func insertQuestions(tx *sql.Tx, questions []questions.Question) error {
	for _, q := range questions {
		fmt.Printf("%s: %s, %s\n", q.Name, q.Part1Answer, q.Part2Answer)
		_, err := tx.Exec(`INSERT INTO questions (
			name, intro, input, part_1_description, part_1_answer, part_2_description, part_2_answer
		) VALUES (?, ?, ?, ?, ?, ?, ?)`, q.Name, q.Intro, q.Input, q.Part1Description, q.Part1Answer, q.Part2Description, q.Part2Answer)
		if err != nil {
			return err
		}
	}
	return nil
}

func handleError(tx *sql.Tx, err error, panicMsg string) {
	fmt.Println(err)
	err = tx.Rollback()
	if err != nil {
		fmt.Println(err)
		panic(panicMsg)
	}
}
