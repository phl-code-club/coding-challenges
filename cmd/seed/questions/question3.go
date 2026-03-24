package questions

import (
	"strconv"
	"strings"
)

var pictogramCount = 1000

type pictogram [5][5]string

func (p *pictogram) String() string {
	var str strings.Builder
	for i, line := range p {
		for _, char := range line {
			str.WriteString(char)
		}
		if i < 4 {
			str.WriteString("\n")
		}
	}
	return str.String()
}

func (p *pictogram) IsValid() bool {
	str := p.String()
	return str == river.String() || str == gem.String() || str == star.String() || str == forest.String()
}

var chars = []string{
	" ",
	"\\",
	"/",
	"|",
	"~",
	"-",
	"_",
	"^",
}

var river = pictogram{
	{" ", "/", "~", "~", "/"},
	{"/", "~", "~", "/", " "},
	{"\\", "~", "~", "\\", " "},
	{" ", "\\", "~", "~", "\\"},
	{" ", "/", "~", "~", "/"},
}

var star = pictogram{
	{"\\", " ", "|", " ", "/"},
	{" ", "\\", "|", "/", " "},
	{"-", "-", "X", "-", "-"},
	{" ", "/", "|", "\\", " "},
	{"/", " ", "|", " ", "\\"},
}

var gem = pictogram{
	{" ", "_", "_", "_", " "},
	{"/", "_", "_", "_", "\\"},
	{"\\", "_", "_", "_", "/"},
	{" ", "\\", "_", "/", " "},
	{" ", " ", "V", " ", " "},
}

var forest = pictogram{
	{" ", " ", "^", " ", " "},
	{" ", "/", "^", "\\", " "},
	{"/", "/", "^", "^", "\\"},
	{" ", "|", "|", "^", "\\"},
	{" ", "|", "|", "|", " "},
}

var pictograms = []pictogram{
	river,
	gem,
	star,
	forest,
}

func generatePictogram() pictogram {
	p := pictogram{}
	if flipCoin(1.0 / 3.0) {
		for i := range 5 {
			line := [5]string{}
			for j := range 5 {
				line[j] = getRandomVal(chars)
			}
			p[i] = line
		}
	} else {
		p = getRandomVal(pictograms)
	}
	return p
}

func generateInput3() Input {
	var input Input
	pictogramList := make([]string, pictogramCount)
	validCount := 0
	longestSeq := 0
	currentSeq := 0
	for i := range pictogramCount {
		p := generatePictogram()
		if p.IsValid() {
			currentSeq++
			if currentSeq > longestSeq {
				longestSeq = currentSeq
			}
			validCount++
		} else {
			currentSeq = 0
		}
		pictogramList[i] = p.String()
	}
	input.Value = strings.Join(pictogramList, "\n")
	input.Part1Answer = strconv.Itoa(validCount * (len(pictogramList) - validCount))
	input.Part2Answer = strconv.Itoa((validCount * (len(pictogramList) - validCount) * longestSeq))
	return input
}

func GenerateQuestion3() Question {
	input := generateInput3()
	return Question{
		Name:             "question3",
		Intro:            readEmbeddedFile("descriptions/question3_intro.md"),
		Input:            input.Value,
		Part1Description: readEmbeddedFile("descriptions/question3_part1.md"),
		Part1Answer:      input.Part1Answer,
		Part2Description: readEmbeddedFile("descriptions/question3_part2.md"),
		Part2Answer:      input.Part2Answer,
		LockCode:         "3",
	}
}
