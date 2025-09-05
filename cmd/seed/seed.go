package main

import (
	"container/heap"
	"database/sql"
	"embed"
	"enchanted-codex/database"
	"flag"
	"fmt"
	"math"
	"math/rand"
	"os"
	"slices"
	"strings"
)

//go:embed descriptions/*.txt
var descriptions embed.FS

var badSpellParts = []string{
	"hex",
	"curse",
	"poison",
	"skull",
	"death",
	"trouble",
	"error",
	"gun",
	"bomb",
	"evil",
}

var spellParts = []string{
	"lizard",
	"fish",
	"leap",
	"chop",
	"kick",
	"punch",
	"block",
	"beam",
	"bean",
	"gas",
	"armor",
	"diaper",
	"energy",
	"smash",
	"crusty",
	"tornado",
	"ritual",
	"blob",
	"bloody",
	"fart",
	"transmute",
	"fireball",
	"charm",
	"spell",
	"enchantment",
	"lucky",
	"chimera",
	"gooey",
	"jelly",
	"smelly",
	"harden",
	"iron",
	"gold",
	"fuzzy",
	"wiggly",
	"bubble",
	"eye",
	"boil",
	"snake",
	"bug",
	"vicious",
	"mockery",
	"wollop",
	"wall",
	"fire",
	"ice",
	"water",
	"air",
	"lightning",
	"burst",
	"wrath",
	"summon",
	"reanimate",
	"fiend",
	"dragon",
	"blast",
	"magic",
	"flesh",
	"tickle",
	"upside down",
	"mirror",
	"world",
	"dimension",
	"byte",
	"slug",
	"assembly",
	"computation",
	"null",
	"tree",
	"earth",
	"spectral",
	"salad",
	"suck",
	"potato",
	"onion",
}

var air = [3]string{
	"~~~",
	"~@~",
	"~~~",
}

var fire = [3]string{
	"@~@",
	"@~@",
	"~@~",
}

var earth = [3]string{
	"~~~",
	"~~~",
	"@@@",
}

var water = [3]string{
	"~@~",
	"~~~",
	"@~@",
}

var wares = []string{
	"herbs",
	"roots",
	"fungi",
	"bones",
	"pelts",
	"organs",
	"elixirs",
	"potions",
	"tonics",
	"books",
	"scrolls",
	"tomes",
	"grimiores",
	"codexes",
	"amulets",
	"rings",
	"circlets",
	"broaches",
	"gems",
	"snacks",
	"refreshments",
	"brooms",
	"pets",
	"hats",
	"maps",
}

var ingredients = []string{
	"mugwort",
	"eye of newt",
	"vinegar",
	"foxglove",
	"cilantro",
	"dragon scale",
	"toad juice",
	"deathclaw saliva",
	"devils cap",
	"kingsfoil",
	"rainbow quartz",
	"mead",
	"ale",
	"philosipher's stone",
	"water",
	"snake venom",
	"unicorn horn",
	"wine",
	"salt",
	"ruby",
	"sapphire",
	"hibiscus",
	"dandelion root",
	"black sand",
	"nettle",
	"peppermint",
	"ash",
	"witchbane",
	"rose petals",
	"mercury",
	"amethyst",
	"scullcap",
	"gryphons wishbone",
	"ectoplasm",
	"coconut milk",
	"sage",
	"crocodile tears",
	"fairy wings",
	"lily of the valley",
	"mountain laurel",
	"vampire dust",
	"marmite",
	"bile",
	"hydrogen peroxide",
	"rubbing alcohol",
}

func main() {
	generateFile := flag.Bool("generate", false, "Generate inputs and answers to files instead of seeding database")
	flag.Parse()

	if *generateFile {
		err := generateToFiles()
		if err != nil {
			panic(fmt.Sprintf("Failed to generate files: %v", err))
		}
		fmt.Println("Generated inputs and answers to files successfully!")
		return
	}

	// Default behavior: seed database
	os.Remove("./test.db")
	db, err := database.GetDB()
	if err != nil {
		panic("No db")
	}

	tx, err := db.Begin()
	if err != nil {
		panic("No tx")
	}

	teamResult, err := tx.Exec("INSERT INTO teams (name) VALUES ('The Enchanted Codex');")
	if err != nil {
		handleError(tx, err, "no team")
		return
	}
	teamID, err := teamResult.LastInsertId()
	if err != nil {
		handleError(tx, err, "no team ID")
		return
	}

	_, err = tx.Exec("INSERT INTO members (name, team_id) VALUES ('Graham', ?), ('Christina', ?);", teamID, teamID)
	if err != nil {
		handleError(tx, err, "no members")
		return
	}

	// Generate and insert questions
	err = insertQuestions(tx)
	if err != nil {
		handleError(tx, err, "no questions")
		return
	}

	tx.Commit()
	fmt.Println("Database seeded successfully!")
}

func generateToFiles() error {
	questions := []Question{
		generateQuestion1(),
		generateQuestion2(),
		generateQuestion3(),
		generateQuestion4(),
		generateQuestion5(),
	}

	// Create output directory if it doesn't exist
	err := os.MkdirAll("generated", 0755)
	if err != nil {
		return fmt.Errorf("failed to create output directory: %v", err)
	}

	for i, q := range questions {
		questionNum := i + 1

		// Write input file
		inputFile := fmt.Sprintf("generated/question%d_input.txt", questionNum)
		err := os.WriteFile(inputFile, []byte(q.Input), 0644)
		if err != nil {
			return fmt.Errorf("failed to write input file %s: %v", inputFile, err)
		}

		// Write answers file
		answersFile := fmt.Sprintf("generated/question%d_answers.txt", questionNum)
		answersContent := fmt.Sprintf("Part 1: %s\nPart 2: %s\n", q.Part1Answer, q.Part2Answer)
		err = os.WriteFile(answersFile, []byte(answersContent), 0644)
		if err != nil {
			return fmt.Errorf("failed to write answers file %s: %v", answersFile, err)
		}

		// Write question details file
		detailsFile := fmt.Sprintf("generated/question%d_details.txt", questionNum)
		detailsContent := fmt.Sprintf("Name: %s\n\nIntro:\n%s\n\nPart 1 Description:\n%s\n\nPart 2 Description:\n%s\n",
			q.Name, q.Intro, q.Part1Description, q.Part2Description)
		err = os.WriteFile(detailsFile, []byte(detailsContent), 0644)
		if err != nil {
			return fmt.Errorf("failed to write details file %s: %v", detailsFile, err)
		}

		fmt.Printf("Generated question %d: %s\n", questionNum, q.Name)
	}

	return nil
}

func insertQuestions(tx *sql.Tx) error {
	questions := []Question{
		generateQuestion1(),
		generateQuestion2(),
		generateQuestion3(),
		generateQuestion4(),
		generateQuestion5(),
	}

	for _, q := range questions {
		_, err := tx.Exec(`INSERT INTO questions (
			name, intro, input, part_1_description, part_1_answer, part_2_description, part_2_answer
		) VALUES (?, ?, ?, ?, ?, ?, ?)`, q.Name, q.Intro, q.Input, q.Part1Description, q.Part1Answer, q.Part2Description, q.Part2Answer)
		if err != nil {
			return err
		}
	}
	return nil
}

type Question struct {
	Name             string
	Intro            string
	Input            string
	Part1Description string
	Part1Answer      string
	Part2Description string
	Part2Answer      string
}

func readEmbeddedFile(filename string) string {
	content, err := descriptions.ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("Failed to read embedded file %s: %v", filename, err))
	}
	return string(content)
}

func generateQuestion1() Question {
	input, totalWords, totalChars, goodWords, goodChars := generateInput1()

	return Question{
		Name:             "Incantation Regulation",
		Intro:            readEmbeddedFile("descriptions/question1_intro.txt"),
		Input:            input,
		Part1Description: readEmbeddedFile("descriptions/question1_part1.txt"),
		Part1Answer:      fmt.Sprintf("%d,%d", totalWords, totalChars),
		Part2Description: readEmbeddedFile("descriptions/question1_part2.txt"),
		Part2Answer:      fmt.Sprintf("%d,%d", goodWords, goodChars),
	}
}

func generateQuestion2() Question {
	input, perfectCounts, brokenCounts := generateInput2()

	return Question{
		Name:             "Cryptarch's Conundrum",
		Intro:            readEmbeddedFile("descriptions/question2_intro.txt"),
		Input:            input,
		Part1Description: readEmbeddedFile("descriptions/question2_part1.txt"),
		Part1Answer:      fmt.Sprintf("%d,%d,%d,%d", perfectCounts[0], perfectCounts[1], perfectCounts[2], perfectCounts[3]),
		Part2Description: readEmbeddedFile("descriptions/question2_part2.txt"),
		Part2Answer:      fmt.Sprintf("%d,%d,%d,%d", perfectCounts[0]+brokenCounts[0], perfectCounts[1]+brokenCounts[1], perfectCounts[2]+brokenCounts[2], perfectCounts[3]+brokenCounts[3]),
	}
}

func generateQuestion3() Question {
	input, maxWare, maxVal, maxTimestamp, maxCount := generateInput3()

	return Question{
		Name:             "Aunt Agnes' Apothecary",
		Intro:            readEmbeddedFile("descriptions/question3_intro.txt"),
		Input:            input,
		Part1Description: readEmbeddedFile("descriptions/question3_part1.txt"),
		Part1Answer:      fmt.Sprintf("%s,%d", maxWare, maxVal),
		Part2Description: readEmbeddedFile("descriptions/question3_part2.txt"),
		Part2Answer:      fmt.Sprintf("%d,%d", maxTimestamp, maxCount),
	}
}

func generateQuestion4() Question {
	input, valid, avgScore := generateInput4()

	return Question{
		Name:             "Potion Commotion",
		Intro:            readEmbeddedFile("descriptions/question4_intro.txt"),
		Input:            input,
		Part1Description: readEmbeddedFile("descriptions/question4_part1.txt"),
		Part1Answer:      fmt.Sprintf("%d", valid),
		Part2Description: readEmbeddedFile("descriptions/question4_part2.txt"),
		Part2Answer:      fmt.Sprintf("%d", avgScore),
	}
}

func generateQuestion5() Question {
	input, validPaths, minPath := generateInput5()

	return Question{
		Name:             "The Mad Mage's Maze",
		Intro:            readEmbeddedFile("descriptions/question5_intro.txt"),
		Input:            input,
		Part1Description: readEmbeddedFile("descriptions/question5_part1.txt"),
		Part1Answer:      fmt.Sprintf("%d", validPaths),
		Part2Description: readEmbeddedFile("descriptions/question5_part2.txt"),
		Part2Answer:      fmt.Sprintf("%d", minPath),
	}
}

func handleError(tx *sql.Tx, err error, panicMsg string) {
	fmt.Println(err)
	err = tx.Rollback()
	if err != nil {
		fmt.Println(err)
		panic(panicMsg)
	}
}

func randBetween(lower, upper int) int {
	if upper < lower {
		panic(fmt.Sprintf("INVALID RANGE %d, %d", lower, upper))
	}
	return rand.Intn(upper-lower) + lower
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

func generateInput1() (string, int, int, int, int) {
	spells := []string{}
	wordCount := 0
	charCount := 0
	badWordCount := 0
	badCharCount := 0
	for range 500 {
		spell := generateSpell()
		spaceCount := strings.Count(spell, " ")
		wordCount += spaceCount + 1
		charCount += len(spell) - spaceCount
		for word := range strings.SplitSeq(spell, " ") {
			if slices.Contains(badSpellParts, word) {
				badWordCount += spaceCount + 1
				badCharCount += len(spell) - spaceCount
				break
			}
		}
		spells = append(spells, spell)
	}
	for i, spell := range spells {
		if strings.Contains(spell, "death") && strings.Contains(spell, "gas") {
			spells[i] = generateSpell()
		}
	}
	return strings.Join(spells, "\n"), wordCount, charCount, wordCount - badWordCount, charCount - badCharCount
}

func generateSpell() string {
	spell := new(strings.Builder)
	numSpellParts := randBetween(2, 5)
	for j := range numSpellParts {
		isBad := flipCoin(1.0 / 5.0)
		if isBad {
			spell.WriteString(getRandomVal(badSpellParts))
		} else {
			spell.WriteString(getRandomVal(spellParts))
		}
		if j < numSpellParts-1 {
			spell.WriteString(" ")
		}
	}
	return spell.String()
}

func generateInput2() (string, [4]int, [4]int) {
	count := [4]int{0, 0, 0, 0}
	brokenCount := [4]int{0, 0, 0, 0}
	elements := [][3]string{air, fire, earth, water}
	runes := [3][]string{}
	for range 500 {
		var rune [3]string
		if flipCoin(1.0 / 3.0) {
			rune = generateRune()
		} else {
			rune = getRandomVal(elements)
		}
		switch rune {
		case air:
			count[0] += 1
		case fire:
			count[1] += 1
		case earth:
			count[2] += 1
		case water:
			count[3] += 1
		}
		broken := checkBrokenRune(rune)
		brokenCount[0] += broken[0]
		brokenCount[1] += broken[1]
		brokenCount[2] += broken[2]
		brokenCount[3] += broken[3]
		runes[0] = append(runes[0], rune[0])
		runes[1] = append(runes[1], rune[1])
		runes[2] = append(runes[2], rune[2])
	}
	runeString := strings.Join([]string{
		strings.Join(runes[0], ""),
		strings.Join(runes[1], ""),
		strings.Join(runes[2], ""),
	}, "\n")
	return runeString, count, brokenCount
}

func checkBrokenRune(rune [3]string) [4]int {
	count := [4]int{0, 0, 0, 0}
	if checkBrokenAir(rune) {
		count[0] += 1
	} else if checkBrokenFire(rune) {
		count[1] += 1
	} else if checkBrokenEarth(rune) {
		count[2] += 1
	} else if checkBrokenWater(rune) {
		count[3] += 1
	}
	return count
}

func checkBrokenAir(rune [3]string) bool {
	return rune[1] == "~~~" && rune[0] == air[0] && rune[2] == air[2]
}

func checkBrokenFire(rune [3]string) bool {
	top := (rune[0] == "~~@" || rune[0] == "@~~") && rune[1] == fire[1] && rune[2] == fire[2]
	middle := (rune[1] == "~~@" || rune[1] == "@~~") && rune[0] == fire[0] && rune[2] == fire[2]
	bottom := rune[2] == "~~~" && rune[0] == fire[0] && rune[1] == fire[1]
	return top || middle || bottom
}

func checkBrokenEarth(rune [3]string) bool {
	return (rune[2] == "@@~" || rune[2] == "@~@" || rune[2] == "~@@") && rune[0] == earth[0] && rune[1] == earth[1]
}

func checkBrokenWater(rune [3]string) bool {
	top := rune[0] == "~~~" && rune[1] == water[1] && rune[2] == water[2]
	bottom := (rune[2] == "@~~" || rune[2] == "~~@") && rune[0] == water[0] && rune[1] == water[1]
	return top || bottom
}

func generateRune() [3]string {
	rune := new([3]string)
	for i := range 3 {
		rune[i] = generateLine()
	}
	return *rune
}

func generateLine() string {
	opts := []string{"~", "@"}
	str := new(strings.Builder)
	for range 3 {
		str.WriteString(getRandomVal(opts))
	}
	return str.String()
}

func generateInput3() (string, string, int, int, int) {
	list := new(strings.Builder)
	time := 0
	timeCounts := [][2]int{}
	currentWindow := [2]int{0, 0}
	wareVals := map[string]int{}
	maxWare := ""
	maxVal := 0
	for range 500 {
		ware := getRandomVal(wares)
		amount := randBetween(1, 999)
		wareVals[ware] = addOrDefault(wareVals, ware, amount)
		if wareVals[ware] > maxVal {
			maxWare = ware
			maxVal = wareVals[ware]
		}
		fmt.Fprintf(list, "%d:%s:%d\n", time, ware, amount)
		currentWindow[1]++
		offset := randBetween(0, 10)
		time += offset
		timeCounts = append(timeCounts, currentWindow)
		currentWindow[0] = time
		currentWindow[1] = 0
	}
	i, j := 0, 1
	maxTimestamp := timeCounts[0][0]
	maxCount := timeCounts[0][1]
	currCount := timeCounts[0][1]
	for {
		if timeCounts[j][0]-100 > timeCounts[i][0] {
			currCount -= timeCounts[i][1]
			i++
		}
		currCount += timeCounts[j][1]
		if currCount > maxCount {
			maxCount = currCount
			maxTimestamp = timeCounts[i][0]
		}
		j++
		if j > len(timeCounts)-1 {
			break
		}
	}
	return list.String(), maxWare, maxVal, maxTimestamp, maxCount
}

func addOrDefault(m map[string]int, key string, defaultValue int) int {
	if value, exists := m[key]; exists {
		return value + defaultValue
	}
	return defaultValue
}

func generateInput4() (string, int, int) {
	recipes := new(strings.Builder)
	valid := 0
	totalScore := 0
	for range 500 {
		total := 0
		max := 0
		ingredientCount := randBetween(1, 10)
		recipeIngredients := map[string]int{}
		for i := range ingredientCount {
			ingredient := getRandomVal(ingredients)
			amount := randBetween(1, 50)
			recipeIngredients[ingredient] = amount
			if amount > max {
				max = amount
			}
			total += amount
			fmt.Fprintf(recipes, "%s:%d", ingredient, amount)
			if i < ingredientCount-1 {
				recipes.WriteString(", ")
			}
		}
		if total >= 100 && max <= total/2 && ingredientCount > 2 {
			score := total
			_, hasKingsfoil := recipeIngredients["kingsfoil"]
			_, hasRainbowQuarts := recipeIngredients["rainbow quartz"]
			foxgloveAmount, hasFoxglove := recipeIngredients["foxglove"]
			mercuryAmount, hasMercury := recipeIngredients["mercury"]
			if hasKingsfoil && hasRainbowQuarts {
				score *= 2
			}
			if hasFoxglove {
				score -= 25 * foxgloveAmount
			}
			if hasMercury {
				score -= 25 * mercuryAmount
			}
			if score < 0 {
				score = 0
			}
			valid++
			totalScore += score
		}
		recipes.WriteString("\n")
	}

	return recipes.String(), valid, totalScore / valid
}

func generateInput5() (string, int, int) {
	maze := new(strings.Builder)
	mazeSlice := [][]string{}
	for range 100 {
		row := []string{}
		for range 100 {
			if flipCoin(1.2 / 4.0) {
				row = append(row, "#")
				maze.WriteString("#")
			} else {
				row = append(row, ".")
				maze.WriteString(".")
			}
		}
		mazeSlice = append(mazeSlice, row)
		maze.WriteString("\n")
	}
	queries := [][4]int{}
	for range 100 {
		startX, startY := randBetween(0, 99), randBetween(0, 99)
		endX, endY := randBetween(0, 99), randBetween(0, 99)
		fmt.Fprintf(maze, "%d %d %d %d\n", startX, startY, endX, endY)
		queries = append(queries, [4]int{startX, startY, endX, endY})
	}
	visited := make([][]bool, len(mazeSlice))
	for i := range visited {
		visited[i] = make([]bool, len(mazeSlice[i]))
	}
	valid := 0
	aStarMin := math.MaxInt
	for _, q := range queries {
		v := slices.Clone(visited)
		dfs([2]int{q[0], q[1]}, mazeSlice, v)
		if v[q[3]][q[2]] {
			valid++
			if dist := aStarPath(mazeSlice, q[0], q[1], q[2], q[3]); dist > 0 && dist < aStarMin {
				aStarMin = dist
			}
		}
	}
	return maze.String(), valid, aStarMin
}

type cell struct {
	reachable bool
	x         int
	y         int
	parent    *cell
	g         int
	h         int
	f         int
}

func newCell(x int, y int, reachable bool) cell {
	return cell{
		reachable: reachable,
		x:         x,
		y:         y,
		parent:    nil,
		g:         0,
		h:         0,
		f:         0,
	}
}

func hasVisited(pos [2]int, visited [][]bool) bool {
	return visited[pos[1]][pos[0]]
}

func isValid(pos [2]int, maze [][]string, visited [][]bool) bool {
	if pos[1] < 0 || pos[1] >= len(maze) {
		return false
	}
	row := maze[pos[1]]
	return pos[0] >= 0 && pos[0] < len(row) && maze[pos[1]][pos[0]] != "#" && !hasVisited(pos, visited)
}

func dfs(start [2]int, maze [][]string, visited [][]bool) {
	if !isValid(start, maze, visited) {
		return
	}
	visited[start[1]][start[0]] = true
	// Left
	dfs([2]int{start[0] - 1, start[1]}, maze, visited)
	// Right
	dfs([2]int{start[0] + 1, start[1]}, maze, visited)
	// Up
	dfs([2]int{start[0], start[1] - 1}, maze, visited)
	// Down
	dfs([2]int{start[0], start[1] + 1}, maze, visited)
}

// Priority queue for A* algorithm
type priorityQueue []*cell

func (pq priorityQueue) Len() int { return len(pq) }

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].f < pq[j].f
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *priorityQueue) Push(x any) {
	item := x.(*cell)
	*pq = append(*pq, item)
}

func (pq *priorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// Manhattan distance heuristic
func manhattanDistance(x1, y1, x2, y2 int) int {
	return int(math.Abs(float64(x1-x2)) + math.Abs(float64(y1-y2)))
}

// A* pathfinding algorithm
func aStarPath(maze [][]string, startX, startY, endX, endY int) int {
	rows := len(maze)
	cols := len(maze[0])

	// Check if start or end positions are walls
	if maze[startY][startX] == "#" || maze[endY][endX] == "#" {
		return -1
	}

	// Check if start and end are the same
	if startX == endX && startY == endY {
		return 0
	}

	// Initialize grid of cells
	grid := make([][]*cell, rows)
	for i := range grid {
		grid[i] = make([]*cell, cols)
		for j := range grid[i] {
			reachable := maze[i][j] != "#"
			c := newCell(j, i, reachable)
			grid[i][j] = &c
		}
	}

	// Priority queue for open set
	openSet := &priorityQueue{}
	heap.Init(openSet)

	// Closed set to track visited cells
	closedSet := make([][]bool, rows)
	for i := range closedSet {
		closedSet[i] = make([]bool, cols)
	}

	// Start cell
	start := grid[startY][startX]
	start.g = 0
	start.h = manhattanDistance(startX, startY, endX, endY)
	start.f = start.g + start.h
	heap.Push(openSet, start)

	// Directions: up, down, left, right
	directions := [][2]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

	for openSet.Len() > 0 {
		current := heap.Pop(openSet).(*cell)

		// If we reached the target
		if current.x == endX && current.y == endY {
			return current.g
		}

		closedSet[current.y][current.x] = true

		// Check all neighbors
		for _, dir := range directions {
			newX := current.x + dir[0]
			newY := current.y + dir[1]

			// Check bounds
			if newX < 0 || newX >= cols || newY < 0 || newY >= rows {
				continue
			}

			neighbor := grid[newY][newX]

			// Skip if wall or already processed
			if !neighbor.reachable || closedSet[newY][newX] {
				continue
			}

			tentativeG := current.g + 1

			// If this path to neighbor is better than any previous one
			if neighbor.parent == nil || tentativeG < neighbor.g {
				neighbor.parent = current
				neighbor.g = tentativeG
				neighbor.h = manhattanDistance(newX, newY, endX, endY)
				neighbor.f = neighbor.g + neighbor.h

				// Add to open set if not already there
				inOpenSet := false
				for _, item := range *openSet {
					if item.x == newX && item.y == newY {
						inOpenSet = true
						break
					}
				}
				if !inOpenSet {
					heap.Push(openSet, neighbor)
				}
			}
		}
	}

	// No path found
	return -1
}
