package questions

import (
	"fmt"
	"strconv"
	"strings"
)

var size = randBetween(800, 1300)

func inBounds(i, j, border int) bool {
	return i <= size-1-border && j <= size-1-border && i >= 0+border && j >= 0+border
}

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

func (m *nodeMap) bfs(n node) int {
	island := 0
	queue := make([]node, 0)
	queue = append(queue, n)
	m.visit(n)
	for len(queue) > 0 {
		island++
		curr := queue[0]
		queue = queue[1:]
		dirs := [4]node{
			newNode(curr.x-1, curr.y),
			newNode(curr.x, curr.y-1),
			newNode(curr.x+1, curr.y),
			newNode(curr.x, curr.y+1),
		}

		for _, d := range dirs {
			if inBounds(d.y, d.x, 0) && !m.hasVisited(d) && m.getNode(d) == '@' {
				queue = append(queue, d)
				m.visit(d)
			}
		}
	}

	return island
}

func solve(m nodeMap) int {
	islandCount := 0
	for y, line := range m.matrix {
		for x, r := range line {
			n := newNode(x, y)
			if r == '@' && !m.hasVisited(n) {
				islandCount++
				m.bfs(n)
			}
		}
	}
	return islandCount * size
}

func solve2(m nodeMap) int {
	islandCount := 0
	islands := make([]int, 0)
	for y, line := range m.matrix {
		for x, r := range line {
			n := newNode(x, y)
			if r == '@' && !m.hasVisited(n) {
				islandCount++
				island := m.bfs(n)
				if island >= 4 {
					islands = append(islands, island)
				}
			}
		}
	}
	totalArea := 0
	for _, island := range islands {
		totalArea += island
	}
	return totalArea * islandCount
}

func generateMatrix() [][]rune {
	m := make([][]rune, size)
	for i := range size {
		m[i] = make([]rune, size)
	}

	for i := range size {
		for j := range size {
			if m[i][j] == '@' {
				continue
			}
			if flipCoin(1.0 / 25.0) {
				m[i][j] = '@'
				if inBounds(i, j, 1) {
					neighbors := [][3]node{
						{newNode(j+1, i-1),
							newNode(j, i-1),
							newNode(j+1, i)},
						{newNode(j+1, i+1),
							newNode(j, i+1),
							newNode(j+1, i)},
						{newNode(j-1, i-1),
							newNode(j, i-1),
							newNode(j-1, i)},
						{newNode(j-1, i+1),
							newNode(j, i+1),
							newNode(j-1, i)},
					}
					for _, n := range neighbors {
						if flipCoin(1.0 / 4.0) {
							a := n[0]
							b := n[1]
							c := n[2]
							m[a.y][a.x] = '@'
							m[b.y][b.x] = '@'
							m[c.y][c.x] = '@'
						}
					}
				}
			} else {
				m[i][j] = '~'
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
	}
	input.Value = matrixToString(matrix)
	input.Part1Answer = strconv.Itoa(solve(m))
	m.trackingMap = make(map[string]bool)
	input.Part2Answer = strconv.Itoa(solve2(m))
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
