package questions

func GenerateQuestion4() Question {
	var input Input
	return Question{
		Name:             "question4",
		Intro:            readEmbeddedFile("descriptions/question4_intro.md"),
		Input:            input.Value,
		Part1Description: readEmbeddedFile("descriptions/question4_part1.md"),
		Part1Answer:      input.Part1Answer,
		Part2Description: readEmbeddedFile("descriptions/question4_part2.md"),
		Part2Answer:      input.Part2Answer,
	}
}
