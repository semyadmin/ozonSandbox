package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func findPossibleFriends(users []map[uint32]struct{}, friends map[uint32]struct{}, user uint32) []uint32 {
	res := make([]uint32, 0, 20)
	var max uint32
	possibleFriends := make(map[uint32]uint32, 20)
	for friend := range friends {
		posFriends := users[friend]
		for posFriend := range posFriends {
			if posFriend == user {
				continue
			}
			if _, ok := friends[posFriend]; ok {
				continue
			}
			possibleFriends[posFriend]++
			if possibleFriends[posFriend] > max {
				max = possibleFriends[posFriend]
			}
		}
	}
	for friend, count := range possibleFriends {
		if count == max {
			res = append(res, friend)
		}
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i] < res[j]
	})
	return res
}

func getResult(users []map[uint32]struct{}, out *bufio.Writer) {
	for i, user := range users {
		if i == 0 {
			continue
		}
		possibleFriends := findPossibleFriends(users, user, uint32(i))
		if len(possibleFriends) == 0 {
			fmt.Fprintf(out, "0")
		}
		for _, friend := range possibleFriends {
			fmt.Fprintf(out, "%d ", friend)
		}
		fmt.Fprintf(out, "\n")
	}
}

// Сканирование данных из консоли
func scan() {
	/* file, _ := os.Open("./tests/12")
	defer file.Close()
	input := bufio.NewReader(file) */
	input := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var n, m, i, u1, u2 uint32
	fmt.Fscan(input, &n, &m)

	users := make([]map[uint32]struct{}, n+1)
	for i = 0; i < n+1; i++ {
		users[i] = make(map[uint32]struct{}, 5)
	}
	for i = 0; i < m; i++ {
		fmt.Fscan(input, &u1, &u2)
		users[u1][u2] = struct{}{}
		users[u2][u1] = struct{}{}
	}
	getResult(users, out)
}

func main() {
	scan()
}
