package questions

import (
	"embed"
	"fmt"
	"math/rand"
)

//go:embed descriptions/*.md
var descriptions embed.FS

type Question struct {
	Name             string
	Intro            string
	Input            string
	Part1Description string
	Part1Answer      string
	Part2Description string
	Part2Answer      string
	LockCode         string
}

type Input struct {
	Value       string
	Part1Answer string
	Part2Answer string
}

func readEmbeddedFile(filename string) string {
	content, err := descriptions.ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("Failed to read embedded file %s: %v", filename, err))
	}
	return string(content)
}

func randBetween(lower, upper int) int {
	if upper < lower {
		panic(fmt.Sprintf("INVALID RANGE %d, %d", lower, upper))
	}
	return rand.Intn(upper-lower) + lower
}

func randFloatBetween(lower, upper float64) float64 {
	if upper < lower {
		panic(fmt.Sprintf("INVALID RANGE %f, %f", lower, upper))
	}
	return lower + rand.Float64()*(upper-lower)
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
	return arr[randBetween(0, size)]
}
