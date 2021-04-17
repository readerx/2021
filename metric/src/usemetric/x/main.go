package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	oscillationPeriod := 10 * time.Minute

	start := time.Now()

	for true {
		oscillationFactor := func() float64 {
			return math.Sin(math.Sin(2 * math.Pi * float64(time.Since(start)) / float64(oscillationPeriod)))
		}
		fmt.Printf("oscillation: %v\n", oscillationFactor())
		time.Sleep(1 * time.Second)
	}
}
