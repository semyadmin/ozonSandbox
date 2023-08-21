package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Time struct {
	slice uint32
	inOut int8
}

func getResult(times []Time) string {
	sort.Slice(times, func(i, j int) bool {
		if times[i].slice == times[j].slice {
			return times[i].inOut > times[j].inOut
		}
		return times[i].slice < times[j].slice
	})
	count := 0
	for i := 0; i < len(times); i++ {
		if count > 1 {
			return "NO"
		}
		if times[i].inOut == 1 {
			count++
		}
		if times[i].inOut == -1 {
			count--
		}
	}

	return "YES"
}

func createTime(t string) (uint32, error) {
	tM := strings.Split(t, ":")
	hour, _ := strconv.Atoi(tM[0])
	if hour > 23 || hour < 0 {
		return 0, fmt.Errorf("invalid hour")
	}
	minute, _ := strconv.Atoi(tM[1])
	if minute > 59 || minute < 0 {
		return 0, fmt.Errorf("invalid minute")
	}
	second, _ := strconv.Atoi(tM[2])
	if second > 59 || second < 0 {
		return 0, fmt.Errorf("invalid second")
	}
	return uint32(hour*3600 + minute*60 + second), nil
}

func parseTime(t string) (uint32, uint32, error) {
	startFinish := strings.Split(t, "-")
	start, err := createTime(startFinish[0])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid time")
	}
	finish, err := createTime(startFinish[1])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid time")
	}
	if finish < start {
		return 0, 0, fmt.Errorf("invalid time")
	}
	return start, finish, nil
}

// Сканирование данных из консоли
func scan() {
	/* file, _ := os.Open("./tests/30")
	defer file.Close()
	input := bufio.NewReader(file) */
	input := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var t, i uint8
	fmt.Fscan(input, &t)
	var n, y uint16
	var dT string
	var start, finish uint32
	var err error
	var no bool
	for i = 0; i < t; i++ {
		fmt.Fscan(input, &n)
		no = false
		slices := make([]Time, 0, n)
		for y = 0; y < n; y++ {
			fmt.Fscan(input, &dT)
			start, finish, err = parseTime(dT)
			if err != nil {
				no = true
			}
			slices = append(slices, Time{slice: start, inOut: 1})
			slices = append(slices, Time{slice: finish, inOut: -1})
		}
		if no {
			fmt.Fprintln(out, "NO")
			continue
		}
		fmt.Fprintln(out, getResult(slices))
	}
}

func main() {
	scan()
}
