package main

import (
	"fmt"
	"math/rand"
	"time"
)

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println("Hello world")
	fmt.Println("Hello world, time now is ", time.Now())
	fmt.Println("Random number is - ", rand.Intn(10))
	fmt.Println(add(1, 2))

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
