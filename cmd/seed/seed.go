package main

import (
	"database/sql"
	"enchanted-codex/database"
	"fmt"
	"math/rand"
	"os"
	"slices"
	"strings"
)

var badSpellParts = []string{
	"hex",
	"curse",
	"poison",
	"skull",
	"death",
	"trouble",
	"error",
	"gun",
	"bomb",
	"evil",
}

var spellParts = []string{
	"lizard",
	"fish",
	"leap",
	"chop",
	"kick",
	"punch",
	"block",
	"beam",
	"bean",
	"gas",
	"armor",
	"diaper",
	"energy",
	"smash",
	"crusty",
	"tornado",
	"ritual",
	"blob",
	"bloody",
	"fart",
	"transmute",
	"fireball",
	"charm",
	"spell",
	"enchantment",
	"lucky",
	"chimera",
	"gooey",
	"jelly",
	"smelly",
	"harden",
	"iron",
	"gold",
	"fuzzy",
	"wiggly",
	"bubble",
	"eye",
	"boil",
	"snake",
	"bug",
	"vicious",
	"mockery",
	"wollop",
	"wall",
	"fire",
	"ice",
	"water",
	"air",
	"lightning",
	"burst",
	"wrath",
	"summon",
	"reanimate",
	"fiend",
	"dragon",
	"blast",
	"magic",
	"flesh",
	"tickle",
	"upside down",
	"mirror",
	"world",
	"dimension",
	"byte",
	"slug",
	"assembly",
	"computation",
	"null",
	"tree",
	"earth",
	"spectral",
	"salad",
	"suck",
	"potato",
	"onion",
}

func main() {
	err := os.Remove("./test.db")
	if err != nil {
		panic("No db file")
	}
	file, err := os.ReadFile("./cmd/seed/questions.sql")
	if err != nil {
		panic("No questions file")
	}
	db, err := database.GetDB()
	if err != nil {
		panic("No db")
	}

	tx, err := db.Begin()
	if err != nil {
		panic("No tx")
	}

	teamResult, err := tx.Exec("INSERT INTO teams (name) VALUES ('The Enchanted Codex');")
	if err != nil {
		handleError(tx, err, "no team")
		return
	}
	teamID, err := teamResult.LastInsertId()
	if err != nil {
		handleError(tx, err, "no team ID")
		return
	}

	_, err = tx.Exec("INSERT INTO members (name, team_id) VALUES ('Graham', ?), ('Christina', ?);", teamID, teamID)
	if err != nil {
		handleError(tx, err, "no team ID")
		return
	}
	_, err = tx.Exec(string(file))
	if err != nil {
		handleError(tx, err, "no questions")
		return
	}

	tx.Commit()
}

func handleError(tx *sql.Tx, err error, panicMsg string) {
	fmt.Println(err)
	err = tx.Rollback()
	if err != nil {
		fmt.Println(err)
		panic(panicMsg)
	}
}

func randBetween(lower, upper int) int {
	if upper < lower {
		panic(fmt.Sprintf("INVALID RANGE %d, %d", lower, upper))
	}
	return rand.Intn(upper-lower) + lower
}

func flipCoin(odds float32) bool {
	return rand.Float32() <= odds
}

func getRandomVal[T any](arr []T) T {
	size := len(arr)
	if size < 1 {
		panic("EMPTY ARR")
	}

	if size == 1 {
		return arr[0]
	}
	return arr[randBetween(0, size-1)]
}

func generateInput1() (string, int, int, int, int) {
	spells := []string{}
	wordCount := 0
	charCount := 0
	badWordCount := 0
	badCharCount := 0
	for range 500 {
		spell := generateSpell()
		spaceCount := strings.Count(spell, " ")
		wordCount += spaceCount + 1
		charCount += len(spell) - spaceCount
		for word := range strings.SplitSeq(spell, " ") {
			if slices.Contains(badSpellParts, word) {
				badWordCount += spaceCount + 1
				badCharCount += len(spell) - spaceCount
				break
			}
		}
		spells = append(spells, spell)
	}
	for i, spell := range spells {
		if strings.Contains(spell, "death") && strings.Contains(spell, "gas") {
			spells[i] = generateSpell()
		}
	}
	return strings.Join(spells, "\n"), wordCount, charCount, wordCount - badWordCount, charCount - badCharCount
}

func generateSpell() string {
	spell := new(strings.Builder)
	numSpellParts := randBetween(2, 5)
	for j := range numSpellParts {
		isBad := flipCoin(1.0 / 5.0)
		if isBad {
			spell.WriteString(getRandomVal(badSpellParts))
		} else {
			spell.WriteString(getRandomVal(spellParts))
		}
		if j < numSpellParts-1 {
			spell.WriteString(" ")
		}
	}
	return spell.String()
}
