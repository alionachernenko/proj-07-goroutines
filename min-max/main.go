package main

import (
	"fmt"
	"math/rand"
	"time"
)

type MinMax struct {
	Min int
	Max int
	in  <-chan []int
	out chan<- MinMax
}

type Generator struct {
	out chan<- []int
	in  <-chan MinMax
}

func main() {
	numbersChan := make(chan []int)
	resultsChan := make(chan MinMax)

	generator := Generator{out: numbersChan, in: resultsChan}
	minmax := MinMax{
		in:  numbersChan,
		out: resultsChan,
	}
	go generator.generateRandomNumbers(10)
	go minmax.findMinMax()

	time.Sleep(2 * time.Second)
}

func (g Generator) generateRandomNumbers(count int) {
	var randomNumbers []int

	for i := 0; i < count; i++ {
		randomNumber := rand.Intn(100)
		randomNumbers = append(randomNumbers, randomNumber)
	}

	g.out <- randomNumbers

	res := <-g.in

	fmt.Printf("Min: %v; Max: %v\n", res.Min, res.Max)
}

func (minmax MinMax) findMinMax() {
	numbers := <-minmax.in

	min := numbers[0]
	max := numbers[0]

	for _, num := range numbers {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	minmax.out <- MinMax{Min: min, Max: max}
}
