package questions

import (
	"fmt"
	"strconv"
	"strings"
)

var latlongCount = 10000

type LatLong struct {
	Lat   float64
	Long  float64
	Name  string
	Valid bool
}

func (l *LatLong) String() string {
	return fmt.Sprintf("(%.4f,%.4f):%s", l.Lat, l.Long, l.Name)
}

var latLongTypes = []string{
	"landmark",
	"clue",
	"trap",
	"thief",
	"merchant",
	"ship",
	"campsite",
	"rumor",
}

func generateLatLong() LatLong {
	lat := randFloatBetween(-90.0, 90.0)
	long := randFloatBetween(-180.0, 180.0)
	name := getRandomVal(latLongTypes)
	return LatLong{
		Lat:   lat,
		Long:  long,
		Name:  name,
		Valid: name == "landmark" || name == "clue",
	}
}

func generateInput1() Input {
	var input Input
	list := make([]string, latlongCount)
	validCount := 0
	landmarkCount := 0
	clueCount := 0
	trapCount := 0
	thiefCount := 0
	for i := range latlongCount {
		val := generateLatLong()
		list[i] = val.String()
		if val.Valid {
			validCount++
		}
		if val.Name == "landmark" {
			landmarkCount++
		}
		if val.Name == "clue" {
			clueCount++
		}
		if val.Name == "trap" {
			trapCount++
		}
		if val.Name == "thief" {
			thiefCount++
		}
	}
	input.Value = strings.Join(list, "\n")
	input.Part1Answer = strconv.Itoa(validCount * 10)
	input.Part2Answer = strconv.Itoa((40 * clueCount) + (30 * landmarkCount) - (10 * trapCount) - (20 * thiefCount))
	return input
}

func GenerateQuestion1() Question {
	input := generateInput1()

	return Question{
		Name:             "LatLong",
		Intro:            readEmbeddedFile("descriptions/question1_intro.md"),
		Input:            input.Value,
		Part1Description: readEmbeddedFile("descriptions/question1_part1.md"),
		Part1Answer:      input.Part1Answer,
		Part2Description: readEmbeddedFile("descriptions/question1_part2.md"),
		Part2Answer:      input.Part2Answer,
	}
}
