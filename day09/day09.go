package day09

import (
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Day struct {

}

func (d *Day) GetDay() string {
	return "day09"
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
	splitted := strings.Split(d.GetInput(), "\n")
	grid := make([][]int, len(splitted))
	for i, line := range splitted {
		grid[i] = make([]int, len(line))
		for j, v := range line {
			grid[i][j], _ = strconv.Atoi(string(v))
		}
	}

	risk := 0
	for i, row := range grid {
		for j := range row {
			if i > 0 && grid[i-1][j] <= grid[i][j] {
				continue
			}
			if i < len(grid) - 1 && grid[i+1][j] <= grid[i][j] {
				continue

			}
			if j > 0 && grid[i][j-1] <= grid[i][j] {
				continue

			}
			if j < len(grid[i]) - 1 && grid[i][j+1] <= grid[i][j] {
				continue
			}

			risk = risk + grid[i][j] + 1
		}
	}
	return strconv.Itoa(risk)
}

func (d *Day) Part2() string {
	splitted := strings.Split(d.GetInput(), "\n")
	grid := make([][]int, len(splitted))
	for i, line := range splitted {
		grid[i] = make([]int, len(line))
		for j, v := range line {
			grid[i][j], _ = strconv.Atoi(string(v))
		}
	}

	basins := make([]int, 0)
	for i, row := range grid {
		for j := range row {
			if i > 0 && grid[i-1][j] <= grid[i][j] {
				continue
			}

			if i < len(grid) - 1 && grid[i+1][j] <= grid[i][j] {
				continue
			}

			if j > 0 && grid[i][j-1] <= grid[i][j] {
				continue
			}

			if j < len(grid[i]) - 1 && grid[i][j+1] <= grid[i][j] {
				continue
			}

			basins = append(basins, countBasins(grid, i, j))
		}
	}

	sort.Ints(basins)
	total := 1
	for _, v := range basins[len(basins)-3:] {
		total = total * v
	}

	return strconv.Itoa(total)
}

func countBasins(grid [][]int, i int, j int) (count int) {
	var neighborns = [4][2]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}

	timesUsed := make(map[[2]int]int)

	pendingToCheck := [][2]int{ {i, j} }
	for len(pendingToCheck) > 0 {
		for _, offset := range neighborns {
			x := pendingToCheck[0][0] + offset[0]
			y := pendingToCheck[0][1] + offset[1]

			timesUsed[[2]int{x, y}]++
			if timesUsed[[2]int{x, y}] > 1 {
				continue
			}

			if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[x]) &&
				grid[x][y] < 9 {
				count++
				pendingToCheck = append(pendingToCheck, [2]int{x, y})
			}
		}

		pendingToCheck = pendingToCheck[1:]
	}

	return count
}
