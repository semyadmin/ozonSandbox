package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Region struct {
	name     string
	inRegion bool
}

var circle = [][2]int{{0, -2}, {-1, -1}, {-1, 1}, {0, 2}, {1, 1}, {1, -1}}

func createRegion(table *[][]Region, region *Region, i, y int) {
	if region.inRegion {
		return
	}
	region.inRegion = true
	for _, c := range circle {
		x := i + c[0]
		y := y + c[1]
		if (*table)[x][y].name == region.name {
			createRegion(table, &(*table)[x][y], x, y)
		}
	}
}

func getResult(table [][]Region) string {
	var region *Region
	checking := make(map[string]struct{}, 5)
	for i := 2; i < len(table)-2; i++ {
		for j := 2; j < len(table[i])-2; j++ {
			region = &table[i][j]
			if region.name == "" {
				continue
			}
			if _, ok := checking[table[i][j].name]; ok {
				if !table[i][j].inRegion {
					return "NO"
				}
				continue
			}
			checking[table[i][j].name] = struct{}{}
			createRegion(&table, &table[i][j], i, j)
		}
	}
	return "YES"
}

// Сканирование данных из консоли
func scan() {
	/* file, _ := os.Open("./tests/11")
	defer file.Close()
	input := bufio.NewReader(file) */
	input := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var t, i uint8
	fmt.Fscan(input, &t)
	var n, m, x, y uint8
	var region string
	var regions []string
	for i = 0; i < t; i++ {
		fmt.Fscan(input, &n, &m)
		table := make([][]Region, n+4)
		table[0] = make([]Region, m+4)
		table[1] = make([]Region, m+4)
		for x = 2; x <= n+1; x++ {
			table[x] = make([]Region, m+4)
			fmt.Fscan(input, &region)
			regions = strings.Split(region, "")
			for y = 0; y < uint8(len(regions)); y++ {
				if regions[y] != "." {
					table[x][y+2].name = regions[y]
				}
			}
		}
		table[n+2] = make([]Region, m+4)
		table[n+3] = make([]Region, m+4)
		fmt.Fprintln(out, getResult(table))
	}
}

func main() {
	scan()
}
