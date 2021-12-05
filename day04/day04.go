package day04

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Day struct {

}

func (d *Day) GetDay() string {
	return "day04"
}

func (d *Day) GetInput() string {
	cnt, _ := ioutil.ReadFile(d.GetDay() + string(os.PathSeparator) + "input.txt")
	return string(cnt)
}

func (d *Day) GetReadme() string {
	cnt, _ := ioutil.ReadFile(d.GetDay() + string(os.PathSeparator) + "readme.MD")
	return string(cnt)
}

func getBoards(input string) []map[string]*Number {
	data := strings.Split(input, "\n")

	boards := make([]map[string]*Number, 0)
	currentBoard := make(map[string]*Number)
	myI := 0
	for i:=2;i< len(data);i++ {
		d := data[i]
		if strings.HasPrefix(d, " ") {
			d = d[1:]
		}
		d = strings.ReplaceAll(d, "  ", " 0")
		for j, v := range strings.Split(d, " ") {
			if vn, err := strconv.Atoi(strings.ReplaceAll(v, "\r", "")); err == nil {
				currentBoard[strconv.Itoa(myI) + "#" + strconv.Itoa(j)] = getNumber(vn)
			}
		}

		if len(data[i]) < 2 {
			myI = -1
			boards = append(boards, currentBoard)
			currentBoard = make(map[string]*Number)
		}
		myI++
	}

	if len(currentBoard) > 0 {
		boards = append(boards, currentBoard)
	}
	return boards
}

func (d *Day) Part1() string {
	data := strings.Split(d.GetInput(), "\n")

	boards := getBoards(d.GetInput())

	for _, v := range strings.Split(data[0], ",") {
		vn, _ := strconv.Atoi(v)
		getNumber(vn).Marked = true

		for i := range boards {
			if isWinner(boards[i]) {
				// winner found!
				return strconv.Itoa(getSumAllUnmarkedNumbers(boards[i]) * vn)
			}
		}
	}

	return ""
}

func (d *Day) Part2() string {
	data := strings.Split(d.GetInput(), "\n")

	boards := getBoards(d.GetInput())

	for _, v := range strings.Split(data[0], ",") {
		vn, _ := strconv.Atoi(v)
		getNumber(vn).Marked = true
		// check for winner

		newBoards := make([]map[string]*Number, 0)
		for i := range boards {
			if !isWinner(boards[i]) {
				newBoards = append(newBoards, boards[i])
			}
		}
		if len(newBoards) == 0 {
			return strconv.Itoa(getSumAllUnmarkedNumbers(boards[0]) * vn)
		}

		boards = newBoards
	}

	return ""
}

func getSumAllUnmarkedNumbers(m map[string]*Number) int {
	val := 0
	for i := 0;i<5;i++ {
		for j := 0;j<5;j++ {
			if v, ok := m[strconv.Itoa(i) + "#" + strconv.Itoa(j)]; ok {
				if !v.Marked {
					val = val + v.Value
				}
			}
		}
	}
	return val
}

func isWinner(m map[string]*Number) bool {
	for i := 0; i <5; i++ {
		winnerX := true
		winnerY := true
		for j:=0;j<5;j++ {
			if !m[strconv.Itoa(i) + "#" + strconv.Itoa(j)].Marked {
				winnerX = false
			}
			if !m[strconv.Itoa(j) + "#" + strconv.Itoa(i)].Marked {
				winnerY = false
			}
		}
		if winnerX {
			return true
		}
		if winnerY {
			return true
		}
	}

	return false
}

var cachedNumbers = make(map[int]*Number)
func getNumber(vn int) *Number {
	if v, ok := cachedNumbers[vn]; ok {
		return v
	} else {
		v := &Number{Value: vn}
		cachedNumbers[vn] = v
		return v
	}
}

type Number struct {
	Value int
	Marked bool
}
