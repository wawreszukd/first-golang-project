package main

import (
	"fmt"
	"math/rand"
	"time"
)

func CreateRandomInt(seed int64, rangeOfRand int) int {
	rand.Seed(seed)
	return rand.Intn(rangeOfRand)
}

func main() {
	fmt.Println(CreateRandomInt(time.Now().UnixNano(), 100))
}
