package day01

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Day struct {

}

func (d *Day) GetDay() string {
	return "day01"
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
	var ctr int
	splitted := strings.Fields(d.GetInput())
	previous, _ := strconv.Atoi(splitted[0])
	for i:=1;i<len(splitted);i++ {
		x, _ := strconv.Atoi(splitted[i])
		if x > previous {
			ctr++
		}
		previous = x
	}

	return strconv.Itoa(ctr)
}

func (d *Day) Part2() string {
	var ctr int
	splitted := strings.Fields(d.GetInput())
	var previous int
	for i:=1;i<len(splitted)-2;i++ {
		x, _ := strconv.Atoi(splitted[i+2])
		previous, _ = strconv.Atoi(splitted[i-1])
		if x > previous {
			ctr++
		}
	}

	return strconv.Itoa(ctr)
}