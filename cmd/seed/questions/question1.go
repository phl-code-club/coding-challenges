package questions

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var latlongCount = 10000

type latLong struct {
	lat   float64
	long  float64
	name  string
	valid bool
}

func (l *latLong) String() string {
	return fmt.Sprintf("(%.4f,%.4f):%s", l.lat, l.long, l.name)
}

func parseLatLong(str string) (latLong, error) {
	rgxp := regexp.MustCompile(`\((\d+\.\d{4}), (\d+\.\d{4})\): (.*)`)
	matches := rgxp.FindStringSubmatch(str)
	if len(matches) < 4 {
		return latLong{}, fmt.Errorf("invalid lat/long line: %s", str)
	}
	lat, err := strconv.ParseFloat(matches[1], 64)
	if err != nil {
		return latLong{}, fmt.Errorf("error parsing latitude %s: %w", matches[1], err)
	}
	if lat < -90.0 || lat > 90.0 {
		return latLong{}, fmt.Errorf("invalid latitude %f", lat)
	}
	long, err := strconv.ParseFloat(matches[2], 64)
	if err != nil {
		return latLong{}, fmt.Errorf("error parsing longitude %s: %w", matches[2], err)
	}
	if long < -180.0 || long > 180.0 {
		return latLong{}, fmt.Errorf("invalid longitude %f", long)
	}
	return latLong{
		lat:  lat,
		long: long,
		name: matches[3],
	}, nil
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

func generateLatLong() latLong {
	lat := randFloatBetween(-90.0, 90.0)
	long := randFloatBetween(-180.0, 180.0)
	name := getRandomVal(latLongTypes)
	return latLong{
		lat:   lat,
		long:  long,
		name:  name,
		valid: name == "landmark" || name == "clue",
	}
}

func generateInput1() Input {
	var input Input
	list := make([]string, latlongCount)
	validCount := 0
	trapCount := 0
	thiefCount := 0
	for i := range latlongCount {
		val := generateLatLong()
		list[i] = val.String()
		if val.valid {
			validCount++
		}
		if val.name == "trap" {
			trapCount++
		}
		if val.name == "thief" {
			thiefCount++
		}
	}
	input.Value = strings.Join(list, "\n")
	input.Part1Answer = strconv.Itoa(validCount)
	input.Part2Answer = strconv.Itoa(validCount - trapCount - (2 * thiefCount))
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
