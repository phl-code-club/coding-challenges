package questions

import (
	"fmt"
	"strconv"
	"strings"
)

var size = 1000

type node struct {
	x int
	y int
}

func (n node) String() string {
	return fmt.Sprintf("%d,%d", n.x, n.y)
}

func newNode(x, y int) node {
	return node{
		x: x,
		y: y,
	}
}

type nodeMap struct {
	matrix      [][]rune
	trackingMap map[string]bool
	islands     int
}

func (m *nodeMap) getNode(n node) rune {
	return m.matrix[n.y][n.x]
}

func (m *nodeMap) hasVisited(n node) bool {
	return m.trackingMap[n.String()]
}

func (m *nodeMap) visit(n node) {
	m.trackingMap[n.String()] = true
}

func (m *nodeMap) bfs(n node) {
	m.islands++
	queue := make([]node, 0)
	queue = append(queue, n)
	m.visit(n)
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		dirs := [4]node{
			newNode(curr.x-1, curr.y),
			newNode(curr.x, curr.y-1),
			newNode(curr.x+1, curr.y),
			newNode(curr.x, curr.y+1),
		}

		for _, d := range dirs {
			if d.x >= 0 && d.x < size && d.y >= 0 && d.y < size && !m.hasVisited(d) && m.getNode(d) == '@' {
				queue = append(queue, d)
				m.visit(d)
			}
		}
	}
}

func solve(m nodeMap) int {
	for y, line := range m.matrix {
		for x, r := range line {
			n := newNode(x, y)
			if r == '@' && !m.hasVisited(n) {
				m.bfs(n)
			}
		}
	}
	return m.islands
}

func generateMatrix() [][]rune {
	m := make([][]rune, size)
	for i := range size {
		m[i] = make([]rune, size)
	}

	for i := range size {
		for j := range size {
			if flipCoin(1.0 / 3.0) {
				m[i][j] = '@'
			} else {
				m[i][j] = '-'
			}
		}
	}

	return m
}

func matrixToString(m [][]rune) string {
	var str strings.Builder
	for i, line := range m {
		for _, r := range line {
			str.WriteRune(r)
		}
		if i < size {
			str.WriteString("\n")
		}
	}
	return str.String()
}

func generateInput4() Input {
	var input Input
	matrix := generateMatrix()
	m := nodeMap{
		matrix:      matrix,
		trackingMap: make(map[string]bool),
		islands:     0,
	}
	input.Value = matrixToString(matrix)
	input.Part1Answer = strconv.Itoa(solve(m))
	return input
}

func GenerateQuestion4() Question {
	input := generateInput4()
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
