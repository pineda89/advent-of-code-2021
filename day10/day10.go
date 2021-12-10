package day10

import (
	"container/list"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Day struct {

}

func (d *Day) GetDay() string {
	return "day10"
}

func (d *Day) GetInput() string {
	cnt, _ := ioutil.ReadFile(d.GetDay() + string(os.PathSeparator) + "input.txt")
	return string(cnt)
}

func (d *Day) GetReadme() string {
	cnt, _ := ioutil.ReadFile(d.GetDay() + string(os.PathSeparator) + "readme.MD")
	return string(cnt)
}

func (d *Day) Part1() string {
	score := map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}

	var total int
	for _, line := range strings.Split(d.GetInput(), "\n") {
		lista := list.New()
		for _, vn := range line {
			v := string(vn)
			if isOpening(v, score) {
				lista.PushBack(v)
			} else if _, ok := score[v]; ok {
				ele := lista.Back()
				if rune(ele.Value.(string)[0]) != rune(v[0]) - 1 &&
					rune(ele.Value.(string)[0]) != rune(v[0]) - 2 {
					total = total + score[v]
				}
				lista.Remove(ele)
			}
		}

	}
	return strconv.Itoa(total)
}

func (d *Day) Part2() string {
	score := map[string]int{
		")": 1,
		"]": 2,
		"}": 3,
		">": 4,
	}

	var scores []int
	for _, line := range strings.Split(d.GetInput(), "\n") {
		lista := list.New()
		validLine := true
		for _, vn := range line {
			v := string(vn)
			if isOpening(v, score) {
				lista.PushBack(v)
			} else if _, ok := score[v]; ok {
				ele := lista.Back()
				if rune(ele.Value.(string)[0]) != rune(v[0]) - 1 &&
					rune(ele.Value.(string)[0]) != rune(v[0]) - 2 {
					validLine = false
				}
				lista.Remove(ele)
			}
		}

		if validLine {
			var tmpScore int
			for lista.Len() > 0 {
				ele := lista.Back()
				tmpScore = (tmpScore * 5) + getScore(ele.Value.(string), score)
				lista.Remove(ele)
			}
			scores = append(scores, tmpScore)
		}
	}

	sort.Ints(scores)
	return strconv.Itoa(scores[len(scores)/2])
}

func isOpening(v string, score map[string]int) bool {
	_, ok := score[ string( rune(v[0]) + 1 ) ]
	_, ok2 := score[ string( rune(v[0]) + 2 ) ]
	return ok || ok2
}

func getScore(v string, score map[string]int) int {
	v1, ok := score[ string( rune(v[0]) + 1 ) ]
	v2, ok2 := score[ string( rune(v[0]) + 2 ) ]
	if ok {
		return v1
	}
	if ok2 {
		return v2
	}
	return 0
}

