package main

import (
	"fmt"
	"math/rand"
	"time"
)

var coinFaces [2]string = [2]string{"HEAD", "TAIL"}

func tossACoin() int {

	nowUnix := time.Now().UnixNano()
	seed := rand.NewSource(nowUnix)
	newRand := rand.New(seed)
	num := newRand.Intn(2)

	// fmt.Printf("num: %v", num)
	// result := coinFaces[num]
	return num
}

func main() {
	trials := 20000
	headCount := 0
	tailCount := 0
	for i := 0; i < trials; i++ {
		result := tossACoin()
		// fmt.Printf("result: %s\n", result)
		if result == 0 {
			headCount++
		} else {
			tailCount++
		}
		// fmt.Printf("result: %s", result)
	}
	fmt.Printf("headCount: %v\n", headCount)
	fmt.Printf("tailCount: %v\n", tailCount)
}
