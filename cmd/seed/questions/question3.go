package questions

func GenerateQuestion3() Question {
	var input Input
	return Question{
		Name:             "question3",
		Intro:            readEmbeddedFile("descriptions/question3_intro.md"),
		Input:            input.Value,
		Part1Description: readEmbeddedFile("descriptions/question3_part1.md"),
		Part1Answer:      input.Part1Answer,
		Part2Description: readEmbeddedFile("descriptions/question3_part2.md"),
		Part2Answer:      input.Part2Answer,
	}
}
