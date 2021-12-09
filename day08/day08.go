package day08

import (
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	ONE_SEGMENTS = 2
	FOUR_SEGMENTS = 4
	SEVEN_SEGMENTS = 3
	EIGHT_SEGMENTS = 7
)

type Day struct {

}

func (d *Day) GetDay() string {
	return "day08"
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
	counter := 0
	for _, line := range strings.Split(d.GetInput(), "\n") {
		output := strings.Fields(strings.Split(line, " | ")[1])
		for _, o := range output {
			if len(o) == ONE_SEGMENTS || len(o) == FOUR_SEGMENTS || len(o) == SEVEN_SEGMENTS || len(o) == EIGHT_SEGMENTS {
				counter++
			}
		}
	}

	return strconv.Itoa(counter)
}

func (d *Day) Part2() string {
	var sum int
	for _, line := range strings.Split(d.GetInput(), "\n") {
		p := strings.Split(line, " | ")
		inputs := strings.Fields(p[0])
		outputs := strings.Fields(p[1])
		wires := make([]string, 10)

		for _, v := range inputs {
			l := len(v)
			switch l {
			case ONE_SEGMENTS: // "the only digit that uses two segments is 1"
				wires[1] = v
			case SEVEN_SEGMENTS: // "Because 7 is the only digit that uses three segments"
				wires[7] = v
			case FOUR_SEGMENTS: // "Because 4 is the only digit that uses four segments"
				wires[4] = v
			case EIGHT_SEGMENTS:
				wires[8] = v // "Because the digits 1, 4, 7, and 8 each use a unique number of segments"
			}
		}

		for i := 0; i < len(inputs); i++ {
			for _, v := range inputs {
				switch len(v) {
				case 5:
					if countOverlaps(v, wires[7]) == 2 && countOverlaps(v, wires[1]) == 1 && countOverlaps(v, wires[3]) == 4 && countOverlaps(v, wires[4]) == 2 {
						wires[2] = v
					}
					if countOverlaps(v, wires[2]) == 3 && countOverlaps(v, wires[4]) == 3 {
						wires[5] = v
					}
					if countOverlaps(v, wires[1]) == 2 {
						wires[3] = v
					}
				case 6:
					if countOverlaps(v, wires[1]) == 1 {
						wires[6] = v
					}
					if countOverlaps(v, wires[1]) == 2 && countOverlaps(v, wires[2]) == 4 && countOverlaps(v, wires[3]) == 4 {
						wires[0] = v
					}
					if countOverlaps(v, wires[7]) == 3 && countOverlaps(v, wires[0]) == 5 {
						wires[9] = v
					}
				}
			}
		}

		var o string
		for k, v := range wires {
			wires[k] = sortStr(v)
		}
		for _, v := range outputs {
			v = sortStr(v)
			for k, w := range wires {
				if w == v {
					o = o + strconv.Itoa(k)
				}
			}
		}
		on, _ := strconv.Atoi(o)
		sum = sum + on
	}
	return strconv.Itoa(sum)
}

func sortStr(input string) string {
	charArray := []rune(input)
	sort.Slice(charArray, func(i int, j int) bool {
		return charArray[i] < charArray[j]
	})
	return string(charArray)
}

func countOverlaps(v, wire string) int {
	count := 0
	for i := range v {
		if strings.Contains(wire, v[i:i+1]) {
			count++
		}
	}
	return count
}