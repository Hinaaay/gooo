package main

import (
	"fmt"
)

var n, size, i, m int
var s []int

type arr struct {
	j []int
}

func (a arr) method(k int) {

	fmt.Println("结果:", a.j)
}

func main() {

	fmt.Scan(&n, &size, "\n")

	if size >= n {
		fmt.Println("It's wrong,try again")
		fmt.Scan(&n, &size, "\n")
	}
	for i = 1; i <= n; i++ {
		fmt.Scan(&m)
		s = append(s, m)
	}
	var h []int
	for i = size - 1; i >= 0; i-- {
		h = append(h, s[i])
	}
	for i = size; i < n; i++ {
		h = append(h, s[i])
	}
	arr1 := arr{
		j: h,
	}
	arr1.method(2)
}
