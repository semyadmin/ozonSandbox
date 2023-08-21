package main

import (
	"bufio"
	"fmt"
	"os"
)

func sum(price, count uint32) uint32 {
	return ((count/3)*2 + count%3) * price
}

func getResult(price map[uint32]uint32) uint32 {
	res := uint32(0)
	for p, v := range price {
		res += sum(p, v)
	}
	return res
}

func scan() {
	/* file, _ := os.Open("./tests/10")
	defer file.Close()
	input := bufio.NewReader(file) */
	input := bufio.NewReader(os.Stdin)
	var t, i uint16

	fmt.Fscan(input, &t)
	var n, p, y uint32

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for i = 0; i < t; i++ {
		fmt.Fscan(input, &n)
		price := make(map[uint32]uint32, n)
		for y = 0; y < n; y++ {
			fmt.Fscan(input, &p)
			price[p]++
		}
		fmt.Fprintln(out, getResult(price))
	}
}

func main() {
	scan()
}
