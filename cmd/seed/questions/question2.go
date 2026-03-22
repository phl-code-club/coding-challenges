package questions

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

func GenerateQuestion2() Question {
	var input Input
	return Question{
		Name:             "question2",
		Intro:            readEmbeddedFile("descriptions/question2_intro.md"),
		Input:            input.Value,
		Part1Description: readEmbeddedFile("descriptions/question2_part1.md"),
		Part1Answer:      input.Part1Answer,
		Part2Description: readEmbeddedFile("descriptions/question2_part2.md"),
		Part2Answer:      input.Part2Answer,
	}
}
