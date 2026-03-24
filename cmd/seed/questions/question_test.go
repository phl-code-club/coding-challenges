package questions_test

import (
	"database/sql"
	"enchanted-codex/cmd/seed/questions"
	"enchanted-codex/database"
	"enchanted-codex/services"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"testing"
	"time"
)

type testCase struct {
	name     string
	question *services.Question
	solver   solver
	want     string
}

type solver func(q *services.Question) (string, error)

var (
	db              *sql.DB
	questionService services.QuestionService
)

func init() {
	d, err := database.GetDB()
	if err != nil {
		panic(err)
	}
	db = d
	questionService = services.NewQuestionService(d)
}

func parseLatLong(str string) (questions.LatLong, error) {
	rgxp := regexp.MustCompile(`\((-?\d+\.\d{4}),(-?\d+\.\d{4})\):(.*)`)
	matches := rgxp.FindStringSubmatch(str)
	if len(matches) < 4 {
		return questions.LatLong{}, fmt.Errorf("invalid lat/long line: %s", str)
	}
	lat, err := strconv.ParseFloat(matches[1], 64)
	if err != nil {
		return questions.LatLong{}, fmt.Errorf("error parsing latitude %s: %w", matches[1], err)
	}
	if lat < -90.0 || lat > 90.0 {
		return questions.LatLong{}, fmt.Errorf("invalid latitude %f", lat)
	}
	long, err := strconv.ParseFloat(matches[2], 64)
	if err != nil {
		return questions.LatLong{}, fmt.Errorf("error parsing longitude %s: %w", matches[2], err)
	}
	if long < -180.0 || long > 180.0 {
		return questions.LatLong{}, fmt.Errorf("invalid longitude %f", long)
	}
	return questions.LatLong{
		Lat:  lat,
		Long: long,
		Name: matches[3],
	}, nil
}

func parseStep(str string) (questions.Step, error) {
	rgxp := regexp.MustCompile(`(\d+): (.+) \(([WNES])\)`)
	matches := rgxp.FindStringSubmatch(str)
	if len(matches) < 4 {
		return questions.Step{}, fmt.Errorf("invalid step line: %s", str)
	}
	secs, err := strconv.ParseInt(matches[1], 10, 0)
	if err != nil {
		return questions.Step{}, fmt.Errorf("invalid timestamp %s: %w", matches[1], err)
	}
	timestamp := time.Unix(secs, 0)
	name := matches[2]
	dir := matches[3]
	return questions.Step{
		Value:     name,
		Direction: dir,
		Timestamp: timestamp,
	}, nil
}

func solveQ1P1(q *services.Question) (string, error) {
	lines := strings.Split(q.Input, "\n")
	validCount := 0
	latLongs := make([]questions.LatLong, len(lines))
	for i, line := range lines {
		latLong, err := parseLatLong(line)
		if err != nil {
			return "", err
		}
		latLongs[i] = latLong
		if latLong.Name == "landmark" || latLong.Name == "clue" {
			validCount++
		}
	}
	return strconv.Itoa(validCount * 10), nil
}

func solveQ1P2(q *services.Question) (string, error) {
	lines := strings.Split(q.Input, "\n")
	landmarkCount := 0
	clueCount := 0
	thiefCount := 0
	trapCount := 0
	latLongs := make([]questions.LatLong, len(lines))
	for i, line := range lines {
		latLong, err := parseLatLong(line)
		if err != nil {
			return "", err
		}
		latLongs[i] = latLong
		switch latLong.Name {
		case "landmark":
			landmarkCount++
		case "clue":
			clueCount++
		case "thief":
			thiefCount++
		case "trap":
			trapCount++
		}
	}
	return strconv.Itoa((40 * clueCount) + (30 * landmarkCount) - (10 * trapCount) - (20 * thiefCount)), nil
}

func solveQ2P1(q *services.Question) (string, error) {
	lines := strings.Split(strings.TrimSpace(q.Input), "\n")
	upDown := 0
	leftRight := 0
	for _, line := range lines {
		step, err := parseStep(line)
		if err != nil {
			return "", err
		}
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
	return strconv.Itoa(upDown * leftRight), nil
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

func solveQ2P2(q *services.Question) (string, error) {
	lines := strings.Split(strings.TrimSpace(q.Input), "\n")
	upDown := 0
	leftRight := 0
	dirs := make([]string, len(lines))
	for i, line := range lines {
		step, err := parseStep(line)
		if err != nil {
			return "", err
		}
		dirs[i] = step.Direction
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
	return strconv.Itoa(upDown*leftRight - loopCount), nil
}

var river = " /~~/\n/~~/ \n\\~~\\ \n \\~~\\\n /~~/"

var star = "\\ | /\n \\|/ \n--X--\n /|\\ \n/ | \\"

var gem = " ___ \n//|\\\\\n\\\\|//\n \\_/ \n  V  "

var forest = "  ^  \n /^\\ \n//^^\\\n ||^\\\n ||| "

var pictogramList = []string{
	river,
	gem,
	star,
	forest,
}

func solveQ3P1(q *services.Question) (string, error) {
	lines := strings.Split(strings.TrimSpace(q.Input), "\n")
	acc := make([]string, 0, 5)
	validCount := 0
	invalidCount := 0
	for _, line := range lines {
		acc = append(acc, line)
		if len(acc) == 5 {
			p := strings.Join(acc, "\n")
			if slices.Contains(pictogramList, p) {
				validCount++
			} else {
				invalidCount++
			}
			acc = make([]string, 0, 5)
		}
	}
	return strconv.Itoa(validCount * invalidCount), nil
}

func solveQ3P2(q *services.Question) (string, error) {
	lines := strings.Split(strings.TrimSpace(q.Input), "\n")
	acc := make([]string, 0, 5)
	validCount := 0
	invalidCount := 0
	currSeq := 0
	longestSeq := 0
	for _, line := range lines {
		acc = append(acc, line)
		if len(acc) == 5 {
			p := strings.Join(acc, "\n")
			if slices.Contains(pictogramList, p) {
				currSeq++
				if currSeq > longestSeq {
					longestSeq = currSeq
				}
				validCount++
			} else {
				currSeq = 0
				invalidCount++
			}
			acc = make([]string, 0, 5)
		}
	}
	return strconv.Itoa(validCount * invalidCount * longestSeq), nil
}

func solveQ4P1(q *services.Question) (string, error) {
	lines := strings.Split(strings.TrimSpace(q.Input), "\n")
	trackingMap := make([][]bool, len(lines))
	islandCount := 0
	islandMap := make([][]string, len(lines))
	queue := make([][2]int, 0)
	for i, line := range lines {
		islandMap[i] = strings.Split(line, "")
		trackingMap[i] = make([]bool, len(line))
	}
	bfs := func(node [2]int) {
		trackingMap[node[1]][node[0]] = true
		queue = append(queue, node)
		for len(queue) > 0 {
			n := queue[0]
			queue = queue[1:]
			directions := [][2]int{
				{
					n[0] - 1,
					n[1],
				},
				{
					n[0] + 1,
					n[1],
				},
				{
					n[0],
					n[1] - 1,
				},
				{
					n[0],
					n[1] + 1,
				},
			}

			for _, d := range directions {
				if d[0] >= 0 &&
					d[0] < len(lines) &&
					d[1] >= 0 &&
					d[1] < len(lines) &&
					!trackingMap[d[1]][d[0]] &&
					islandMap[d[1]][d[0]] == "@" {
					queue = append(queue, d)
					trackingMap[d[1]][d[0]] = true
				}
			}

		}
	}
	for i, line := range islandMap {
		for j, char := range line {
			if char == "@" && !trackingMap[i][j] {
				islandCount++
				bfs([2]int{j, i})
			}
		}
	}
	return strconv.Itoa(islandCount * len(lines)), nil
}
func solveQ4P2(q *services.Question) (string, error) {
	lines := strings.Split(strings.TrimSpace(q.Input), "\n")
	trackingMap := make([][]bool, len(lines))
	islandCount := 0
	islandMap := make([][]string, len(lines))
	queue := make([][2]int, 0)
	for i, line := range lines {
		islandMap[i] = strings.Split(line, "")
		trackingMap[i] = make([]bool, len(line))
	}
	bfs := func(node [2]int) int {
		islandSize := 0
		trackingMap[node[1]][node[0]] = true
		queue = append(queue, node)
		for len(queue) > 0 {
			islandSize++
			n := queue[0]
			queue = queue[1:]
			directions := [][2]int{
				{
					n[0] - 1,
					n[1],
				},
				{
					n[0] + 1,
					n[1],
				},
				{
					n[0],
					n[1] - 1,
				},
				{
					n[0],
					n[1] + 1,
				},
			}

			for _, d := range directions {
				if d[0] >= 0 &&
					d[0] < len(lines) &&
					d[1] >= 0 &&
					d[1] < len(lines) &&
					!trackingMap[d[1]][d[0]] &&
					islandMap[d[1]][d[0]] == "@" {
					queue = append(queue, d)
					trackingMap[d[1]][d[0]] = true
				}
			}
		}

		return islandSize
	}

	islandArea := 0
	for i, line := range islandMap {
		for j, char := range line {
			if char == "@" && !trackingMap[i][j] {
				islandCount++
				area := bfs([2]int{j, i})
				if area >= 4 {
					islandArea += area
				}
			}
		}
	}
	return strconv.Itoa(islandCount * islandArea), nil
}

var solvers = [][2]solver{
	{
		solveQ1P1,
		solveQ1P2,
	},
	{
		solveQ2P1,
		solveQ2P2,
	},
	{
		solveQ3P1,
		solveQ3P2,
	},
	{
		solveQ4P1,
		solveQ4P2,
	},
}

func TestQuestion1(t *testing.T) {
	qs, err := questionService.ListQuestions()
	if err != nil {
		t.Errorf("Error getting questions: %v", err)
	}
	tests := make([]testCase, 0)
	for i, q := range qs {
		question, _ := questionService.GetQuestionByID(q.ID)
		tc1 := testCase{
			name:     question.Name + " Part 1",
			question: question,
			solver:   solvers[i][0],
			want:     question.Part1Answer,
		}
		tc2 := testCase{
			name:     question.Name + " Part 2",
			question: question,
			solver:   solvers[i][1],
			want:     question.Part2Answer,
		}
		tests = append(tests, tc1, tc2)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.solver(tt.question)
			if err != nil {
				t.Errorf("Failed parsing input: %v", err)
			}
			if tt.want != got {
				t.Errorf("GenerateQuestion1() = %v, want %v", got, tt.want)
			}
		})
	}

}
