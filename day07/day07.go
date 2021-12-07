package day07

import (
	"io/ioutil"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Day struct {

}

func (d *Day) GetDay() string {
	return "day07"
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
	crabs := make([]int, 0)
	for _, v := range strings.Split(d.GetInput(), ",") {
		vn, _ := strconv.Atoi(v)
		crabs = append(crabs, vn)
	}

	med := median(crabs)

	counter := 0
	for _, v := range crabs {
		counter = counter + int(math.Abs(float64(v - med)))
	}

	return strconv.Itoa(counter)
}

func (d *Day) Part2() string {
	crabs := make([]int, 0)
	for _, v := range strings.Split(d.GetInput(), ",") {
		vn, _ := strconv.Atoi(v)
		crabs = append(crabs, vn)
	}

	mean := mean(crabs)

	var underMean, overMean int
	for _, v := range crabs {
		n := math.Abs(float64(v - mean))
		underMean = underMean + calculateCost(int(n))

		n = math.Abs(float64(v - mean + 1))
		overMean = overMean + calculateCost(int(n))
	}

	return strconv.Itoa(int(math.Min(float64(underMean), float64(overMean))))
}

func calculateCost(n int) int {
	return n * (n + 1) / 2
}

func median(input []int) (median int) {
	if len(input) != 0 {
		if !sort.IntsAreSorted(input) {
			sort.Ints(input)
		}

		if len(input)%2 == 0 {
			median = mean(input[len(input)/2-1 : len(input)/2+1])
		} else {
			median = input[len(input)/2]
		}
	}
	return
}

func mean(ints []int) int {
	if len(ints) != 0 {
		sum := 0
		for _, v := range ints {
			sum = sum + v
		}
		return sum / len(ints)
	}
	return 0
}