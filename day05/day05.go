package day05

import (
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

type Day struct {

}

func (d *Day) GetDay() string {
	return "day05"
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
	mapa := make(map[string]int)
	for _, v := range strings.Split(d.GetInput(), "\n") {
		x1, y1, x2, y2 := getCoordinates(v)

		directionX := getDirection(x1, x2)
		directionY := getDirection(y1, y2)

		for i:=x1;i!=x2+directionX;i=i+directionX {
			for j:=y1;j!=y2+directionY;j=j+directionY {
				if x1 == x2 || y1 == y2 {
					mapa[strconv.Itoa(i) + "#" + strconv.Itoa(j)]++
				}
			}
		}
	}

	ctr := 0
	for _, v := range mapa {
		if v >= 2 {
			ctr++
		}
	}
	return strconv.Itoa(ctr)
}

func (d *Day) Part2() string {
	mapa := make(map[string]int)
	for _, v := range strings.Split(d.GetInput(), "\n") {
		x1, y1, x2, y2 := getCoordinates(v)

		directionX := getDirection(x1, x2)
		directionY := getDirection(y1, y2)

		for i:=x1;i!=x2+directionX;i=i+directionX {
			for j:=y1;j!=y2+directionY;j=j+directionY {
				if x1 == x2 || y1 == y2 {
					mapa[strconv.Itoa(i) + "#" + strconv.Itoa(j)]++
				} else {
					// diagonal
					if int(math.Abs(float64(i-x1))) == int(math.Abs(float64(j-y1))) {
						mapa[strconv.Itoa(i) + "#" + strconv.Itoa(j)]++
					}
				}
			}
		}
	}

	ctr := 0
	for _, v := range mapa {
		if v >= 2 {
			ctr++
		}
	}
	return strconv.Itoa(ctr)
}

func getCoordinates(v string) (int, int, int, int) {
	v = strings.ReplaceAll(v, "\r", "")
	s := strings.Split(v, " -> ")
	t := strings.Split(s[0], ",")
	x1, _ := strconv.Atoi(t[0])
	y1, _ := strconv.Atoi(t[1])
	t = strings.Split(s[1], ",")
	x2, _ := strconv.Atoi(t[0])
	y2, _ := strconv.Atoi(t[1])
	return x1, y1, x2, y2
}

func getDirection(x1 int, x2 int) int {
	if x2 >= x1 {
		return 1
	}
	return -1
}
