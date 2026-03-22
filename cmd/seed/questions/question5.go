package questions

func GenerateQuestion5() Question {
	var input Input
	return Question{
		Name:             "question5",
		Intro:            readEmbeddedFile("descriptions/question5_intro.md"),
		Input:            input.Value,
		Part1Description: readEmbeddedFile("descriptions/question5_part1.md"),
		Part1Answer:      input.Part1Answer,
		Part2Description: readEmbeddedFile("descriptions/question5_part2.md"),
		Part2Answer:      input.Part2Answer,
	}
}
