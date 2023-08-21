package main

import (
	"bufio"
	"fmt"
	"os"
)

type Employer struct {
	level  uint8
	id     uint8
	inPear bool
}

func abs(x, y uint8) uint8 {
	if x > y {
		x, y = y, x
	}
	return y - x
}

func getResult(emp []uint8, out *bufio.Writer) {
	findEmployer := make(map[uint8]struct{}, len(emp))
	var secondEmp, i, y, e uint8

	for i = 0; i < uint8(len(emp)); i++ {
		e = emp[i]
		secondEmp = 0
		if _, ok := findEmployer[i]; ok {
			continue
		}
		for y = i + 1; y < uint8(len(emp)); y++ {
			if _, ok := findEmployer[y]; ok {
				continue
			}
			if secondEmp == 0 {
				secondEmp = y
			}
			if abs(emp[y], e) < abs(emp[secondEmp], e) {
				secondEmp = y
			}
		}
		findEmployer[i] = struct{}{}
		findEmployer[secondEmp] = struct{}{}
		fmt.Fprintf(out, "%d %d\n", i+1, secondEmp+1)
	}
}

func scan() {
	/* file, _ := os.Open("./tests/20")
	defer file.Close()
	input := bufio.NewReader(file) */
	input := bufio.NewReader(os.Stdin)
	var t uint8
	fmt.Fscan(input, &t)
	var n, a, i, y uint8
	var employers []uint8
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for i = 0; i < t; i++ {
		fmt.Fscan(input, &n)
		employers = make([]uint8, n)
		for y = 0; y < n; y++ {
			fmt.Fscan(input, &a)
			employers[y] = a
		}
		getResult(employers, out)

		fmt.Fprintln(out, "")
	}
}

func main() {
	scan()
}
