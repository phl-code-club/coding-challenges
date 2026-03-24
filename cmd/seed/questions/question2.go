package questions

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var stepCount = 10000

var dayWindow = 80

const DayInSeconds = 60 * 60 * 24

var clues = []string{
	"footprints",
	"disturbed earth",
	"torn fabric",
	"ditched supplies",
	"bones",
	"marked trees",
}

var landmarks = []string{
	"spooky trees",
	"jagged rocks",
	"quicksand",
	"tall cliffs",
	"river rapids",
	"waterfall",
	"dark ravine",
	"creepy village",
	"abandoned fort",
	"winding road",
}

var directions = []string{
	"N",
	"S",
	"E",
	"W",
}

var loops = map[string]struct{}{
	"NWSE": {},
	"NESW": {},
	"SWNE": {},
	"SENW": {},
	"WNES": {},
	"WSEN": {},
	"ESWN": {},
	"ENWS": {},
}

type Step struct {
	Value     string
	Direction string
	Timestamp time.Time
}

func (s *Step) String() string {
	return fmt.Sprintf("%d: %s (%s)", s.Timestamp.Unix(), s.Value, s.Direction)
}

func randomTimestamp() time.Time {
	day := rand.Int63n(int64(DayInSeconds * dayWindow))
	return time.Unix(time.Now().Unix()-day, 0)
}

func randomStep() Step {
	timestamp := randomTimestamp()
	direction := getRandomVal(directions)
	var value string
	if flipCoin(1.0 / 2.0) {
		value = getRandomVal(landmarks)
	} else {
		value = getRandomVal(clues)
	}
	return Step{
		Value:     value,
		Direction: direction,
		Timestamp: timestamp,
	}
}

func generateInput2() Input {
	var input Input
	steps := make([]Step, stepCount)
	upDown := 0
	leftRight := 0
	dirs := ""
	for i := range stepCount {
		step := randomStep()
		dirs = dirs + step.Direction
		steps[i] = step
		input.Value += step.String() + "\n"
		switch step.Direction {
		case "N":
			upDown += 10
		case "S":
			upDown -= 10
		case "E":
			leftRight += 10
		case "W":
			leftRight -= 10
		}
	}
	loopCount := 0
	for i := 0; i < len(dirs)-3; i++ {
		set := string(dirs[i]) + string(dirs[i+1]) + string(dirs[i+2]) + string(dirs[i+3])
		if _, ok := loops[set]; ok {
			loopCount++
			i += 3
		}
	}

	input.Part1Answer = strconv.Itoa(upDown * leftRight)
	input.Part2Answer = strconv.Itoa(upDown*leftRight - loopCount)
	return input
}

func GenerateQuestion2() Question {
	input := generateInput2()
	return Question{
		Name:             "A Perilous Journey",
		Intro:            readEmbeddedFile("descriptions/question2_intro.md"),
		Input:            input.Value,
		Part1Description: readEmbeddedFile("descriptions/question2_part1.md"),
		Part1Answer:      input.Part1Answer,
		Part2Description: readEmbeddedFile("descriptions/question2_part2.md"),
		Part2Answer:      input.Part2Answer,
	}
}
