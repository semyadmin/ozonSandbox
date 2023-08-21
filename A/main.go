package main

import (
	"bufio"
	"fmt"
	"os"
)

func sum(a int16, b int16) int16 {
	return a + b
}

func scan() {
	/* file, _ := os.Open("./test/test.txt")
	defer file.Close()
	in := bufio.NewReader(file) */
	in := bufio.NewReader(os.Stdin)
	var a int16 // переменная a
	var b int16 // переменная b
	var t int   // количество итераций считывания данных из консоли

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	fmt.Fscan(in, &t)

	for i := 0; i < t; i++ {
		fmt.Fscan(in, &a, &b)
		fmt.Fprintln(out, sum(a, b))
	}
}

func main() {
	// Данные для подсчета
	scan()
}
