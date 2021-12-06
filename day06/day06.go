package day06

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Day struct {

}

func (d *Day) GetDay() string {
	return "day06"
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
	// of course, the part2 is also the solution for part1. But this was my first proposal.
	values := toIntArray(strings.Split(d.GetInput(), ","))
	for iterations := 0;iterations < 80;iterations++ {
		newValues := make([]int, 0)
		for i := range values {
			v := values[i] - 1
			if v < 0 {
				v = 6
				newValues = append(newValues, 8)
			}
			newValues = append(newValues, v)
		}
		values = newValues
	}

	return strconv.Itoa(len(values))
}

func (d *Day) Part2() string {
	values := toIntArray(strings.Split(d.GetInput(), ","))
	iterations := 256
	week := make([]int, 9)
	for i := range values {
		week[values[i]]++
	}

	for day:=0;day<iterations;day++ {
		today := week[0]
		for daysToNewBorn:=0; daysToNewBorn<len(week)-1;daysToNewBorn++ {
			week[daysToNewBorn] = week[daysToNewBorn+1]
		}
		week[6], week[8] = week[6] + today, today
	}

	sum := 0
	for i := range week {
		sum += week[i]
	}

	return strconv.Itoa(sum)
}

func toIntArray(split []string) []int {
	result := make([]int, len(split))
	for i := range split {
		result[i], _ = strconv.Atoi(split[i])
	}
	return result
}