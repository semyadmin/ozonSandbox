package main

import (
	"bufio"
	"fmt"
	"os"
)

func getResult(dict map[uint16]struct{}, report []uint16) string {
	dict[report[0]] = struct{}{}
	for i := 1; i < len(report); i++ {
		if _, ok := dict[report[i]]; ok && report[i] != report[i-1] {
			return "NO"
		}
		dict[report[i]] = struct{}{}
	}
	return "YES"
}

func scan() {
	/* file, _ := os.Open("./tests/07")
	defer file.Close()
	input := bufio.NewReader(file) */
	input := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var t, i uint8
	var n, j uint16
	fmt.Fscan(input, &t)
	for i = 0; i < t; i++ {
		fmt.Fscan(input, &n)
		dict := make(map[uint16]struct{}, n)
		report := make([]uint16, n)
		for j = 0; j < n; j++ {
			fmt.Fscan(input, &report[j])
		}
		fmt.Fprintf(out, "%s\n", getResult(dict, report))
	}
}

func main() {
	scan()
}
