package day03

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type FILTERTYPE int
const FILTERTYPE_POPULAR FILTERTYPE = 1
const FILTERTYPE_UNPOPULAR FILTERTYPE = -1

type Day struct {

}

func (d *Day) GetDay() string {
	return "day03"
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
	data := strings.Split(d.GetInput(), "\n")
	var gamma string
	for _, v := range getAllOnesAppearences(data) {
		gamma += filter(FILTERTYPE_POPULAR, v, len(data))
	}

	numericGamma, _ := strconv.ParseUint(gamma, 2, 64)
	epsilon := strconv.FormatUint(^numericGamma, 2)
	numericEpsilon, _ := strconv.ParseUint(epsilon[len(epsilon) - len(data[0]):], 2, 64)

	return strconv.FormatUint(numericGamma * numericEpsilon, 10)
}

func (d *Day) Part2() string {
	oxygen := applyBitCriteria(strings.Split(d.GetInput(), "\n"), FILTERTYPE_POPULAR)
	co2 := applyBitCriteria(strings.Split(d.GetInput(), "\n"), FILTERTYPE_UNPOPULAR)

	oxygenNum, _ := strconv.ParseInt(oxygen, 2, 64)
	co2Num, _ := strconv.ParseInt(co2, 2, 64)

	return strconv.FormatInt(oxygenNum * co2Num, 10)
}

func getAllOnesAppearences(input []string) []int {
	var columns = make([]int, len(input[0]))
	for _, v := range input {
		for j := range v {
			if v[j:j+1] == "1" {
				columns[j]++
			}
		}
	}
	return columns
}

func applyBitCriteria(split []string, FILTERTYPE FILTERTYPE) string {
	filtered := split
	var i = 0
	for len(filtered) > 1 {
		filteredVal := filter(FILTERTYPE, getAllOnesAppearences(filtered)[i], len(filtered))

		newFiltered := make([]string, 0)
		for _, line := range filtered {
			if line[i:i+1] == filteredVal {
				newFiltered = append(newFiltered, line)
			}
		}
		filtered = newFiltered
		i++
	}
	return filtered[0]
}

func filter(filtertype FILTERTYPE, quantity int, totalRows int) string {
	if filtertype == FILTERTYPE_POPULAR {
		if totalRows-quantity <= quantity {
			return "1"
		} else {
			return "0"
		}
	} else {
		if totalRows-quantity > quantity {
			return "1"
		} else {
			return "0"
		}
	}
}