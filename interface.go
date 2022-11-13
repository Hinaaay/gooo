package main

import (
	"fmt"
	"sort"
)

type sum []string

func (s sum) Len() int {
	return len(s)
}

func (s sum) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sum) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	nums := sum{
		"6",
		"5",
		"8",
		"3",
		"9",
		"0",
		"1",
		"2",
	}
	sort.Sort(nums)
	for _, v := range nums {
		fmt.Printf("%s\n", v)
	}
}
