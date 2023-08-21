package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Processor struct {
	energyInSecond uint64
	freeTime       uint64
	id             uint32
	position       uint32
}

type Time struct {
	in       uint64
	duration uint64
}

func getResult(processors []*Processor, timeArr []*Time, timeZone []uint64) uint64 {
	var res uint64
	var indexTimeZone uint64
	timeZoneMap := make(map[uint64][]*Processor, len(processors))
	var proc *Processor
	for _, t := range timeArr {
		for timeZone[indexTimeZone] <= t.in {
			if _, ok := timeZoneMap[timeZone[indexTimeZone]]; ok {
				for _, p := range timeZoneMap[timeZone[indexTimeZone]] {
					addToHeap(&processors, p)
				}
				delete(timeZoneMap, timeZone[indexTimeZone])
			}
			indexTimeZone++
		}
		if len(processors) > 0 {
			proc = pop(&processors)
			res += proc.energyInSecond * t.duration
			timeZoneMap[t.duration+t.in] = append(timeZoneMap[t.duration+t.in], proc)
		}
	}

	return res
}

func pop(heap *[]*Processor) *Processor {
	proc := (*heap)[0]
	(*heap)[0] = (*heap)[len(*heap)-1]
	var i uint32
	var child uint32
	for 2*i+2 < uint32(len(*heap)) {
		child = 2*i + 1
		if (*heap)[child].energyInSecond > (*heap)[2*i+2].energyInSecond {
			child = 2*i + 2
		}
		if (*heap)[i].energyInSecond > (*heap)[child].energyInSecond {
			(*heap)[i], (*heap)[child] = (*heap)[child], (*heap)[i]
			i = child
			continue
		}
		break
	}
	(*heap) = (*heap)[:len(*heap)-1]
	return proc
}

func addToHeap(heap *[]*Processor, processor *Processor) {
	(*heap) = append((*heap), processor)
	var i uint32
	i = uint32(len(*heap) - 1)
	for i > 0 && (*heap)[(i-1)/2].energyInSecond > (*heap)[i].energyInSecond {
		(*heap)[(i-1)/2], (*heap)[i] = (*heap)[i], (*heap)[(i-1)/2]
		i = (i - 1) / 2
	}
}

// Сканирование данных из консоли
func scan() {
	/* file, _ := os.Open("./tests/01")
	defer file.Close()
	input := bufio.NewReader(file) */
	input := bufio.NewReader(os.Stdin)
	var n, m uint32
	fmt.Fscan(input, &n, &m)
	processors := make([]*Processor, 0, n)
	var i uint32
	for i = 0; i < n; i++ {
		proc := &Processor{id: i}
		fmt.Fscan(input, &proc.energyInSecond)
		addToHeap(&processors, proc)
	}
	var t, l uint64
	timeArr := make([]*Time, m)
	timeZones := make([]uint64, m)
	for i = 0; i < m; i++ {
		fmt.Fscan(input, &t, &l)
		timeArr[i] = &Time{in: t, duration: l}
		timeZones[i] = t + l
	}
	sort.Slice(timeZones, func(i, j int) bool {
		return timeZones[i] < timeZones[j]
	})
	fmt.Println(getResult(processors, timeArr, timeZones))
}

func main() {
	scan()
}
