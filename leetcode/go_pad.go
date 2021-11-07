package main

import (
	"sort"
)

func driver(input [][]int) int {
	return maxTwoEvents(input)
}

func maxTwoEvents(events [][]int) int {
	sort.Slice(events, func(i, j int) bool {
		return events[i][0] < events[j][0]
	}) // sort by start time

	eventsQueue := make([][]int, len(events))
	copy(eventsQueue, events)

	sort.Slice(eventsQueue, func(i, j int) bool {
		return eventsQueue[i][1] < eventsQueue[j][1]
	}) // sort by end time

	ans := 0
	firstEvent := 0

	for _, event := range events {
		for len(eventsQueue) != 0 && event[0] > eventsQueue[0][1] {
			firstEvent = getMax(firstEvent, eventsQueue[0][2])
			eventsQueue = eventsQueue[1:]
		}
		ans = getMax(ans, firstEvent+event[2])
	}

	return ans
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
