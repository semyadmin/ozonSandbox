package main

import (
	"bufio"
	"fmt"
	"os"
)

func getResult(table [][]uint8, clicks []uint8, out *bufio.Writer) {
	// Применяем пузырьковую сортировку
	for i := 0; i < len(clicks); i++ {
		p := clicks[i] - 1
		for y := 0; y < len(table[0])-1; y++ {
			for x := 0; x < (len(table[0])-1)-y; x++ {
				if table[p][x] > table[p][x+1] {
					for z := 0; z < len(table); z++ {
						table[z][x], table[z][x+1] = table[z][x+1], table[z][x]
					}
				}
			}
		}
	}
	for i := 0; i < len(table[0]); i++ {
		for j := 0; j < len(table); j++ {
			fmt.Fprint(out, table[j][i], " ")
		}
		fmt.Fprintln(out)
	}
}

func scan() {
	/* file, _ := os.Open("./tests/10")
	defer file.Close()
	input := bufio.NewReader(file) */
	input := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var t, n, m, i, y, x, k uint8
	fmt.Fscan(input, &t)
	for i = 0; i < t; i++ {
		fmt.Fscan(input, &n, &m)
		table := make([][]uint8, m)
		for y = 0; y < m; y++ {
			table[y] = make([]uint8, n)
		}
		for y = 0; y < n; y++ {
			for x = 0; x < m; x++ {
				fmt.Fscan(input, &table[x][y])
			}
		}
		fmt.Fscan(input, &k)
		clicks := make([]uint8, k)
		for y = 0; y < k; y++ {
			fmt.Fscan(input, &clicks[y])
		}
		getResult(table, clicks, out)
		fmt.Fprintln(out)
	}
}

func main() {
	scan()
}
