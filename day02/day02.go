package day02

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Day struct {

}

func (d *Day) GetDay() string {
	return "day02"
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
	var horizontal, depth int
	for _, v := range strings.Split(d.GetInput(), "\n") {
		vs := strings.Split(v, " ")
		val, _ := strconv.Atoi(vs[1])
		switch vs[0] {
		case "forward":
			horizontal += val
		case "down":
			depth += val
		case "up":
			depth -= val
		}
	}

	return strconv.Itoa(horizontal * depth)
}

func (d *Day) Part2() string {
	var horizontal, depth, aim int
	for _, v := range strings.Split(d.GetInput(), "\n") {
		vs := strings.Split(v, " ")
		val, _ := strconv.Atoi(vs[1])
		switch vs[0] {
		case "forward":
			horizontal += val
			depth += aim * val
		case "down":
			aim += val
		case "up":
			aim -= val
		}
	}

	return strconv.Itoa(horizontal * depth)
}