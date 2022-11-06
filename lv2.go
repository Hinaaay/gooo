package main

import (
	"fmt"
	"math/rand"
	"time"
)

var h []int32

func main() {
	min := int32(1)
	max := int32(500)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		we := rand.Int31n(max-min) + min
		h = append(h, we)
	}
	fmt.Printf("Input slice: %v\n", h)
}
