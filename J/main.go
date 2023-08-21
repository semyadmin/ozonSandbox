package main

import (
	"bufio"
	"fmt"
	"os"
)

func getResult(dict *map[string][]string, str string) string {
	var res string
	i := 0
	for i < len(str) {
		if res != "" {
			break
		}
		s, ok := (*dict)[str[i:]]
		if ok {
			for _, st := range s {
				if st != str && len(st) > len(res) {
					res = st
				}
			}
		}
		i++
	}
	for _, s := range *dict {
		if res != "" {
			break
		}
		for _, st := range s {
			if st != str {
				res = st
				break
			}
		}
	}

	return res
}

func addToDictionary(dict *map[string][]string, str string) {
	i := 0
	for i < len(str) {
		(*dict)[str[i:]] = append((*dict)[str[i:]], str)
		i++
	}
}

// Сканирование данных из консоли
func scan() {
	/* 	file, _ := os.Open("./tests/05")
	defer file.Close()
	   	input := bufio.NewReader(file) */
	input := bufio.NewReader(os.Stdin)
	var n, i uint16
	fmt.Fscan(input, &n)
	var str string
	dict := make(map[string][]string, n*10)
	for i = 0; i < n; i++ {
		fmt.Fscan(input, &str)
		addToDictionary(&dict, str)
	}
	var q uint16
	fmt.Fscan(input, &q)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for i = 0; i < q; i++ {
		fmt.Fscan(input, &str)
		fmt.Fprintln(out, getResult(&dict, str))
	}
}

func main() {
	scan()
}
