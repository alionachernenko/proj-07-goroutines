package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Generator struct {
	out chan<- []int
}

type Calculator struct {
	in  <-chan []int
	out chan<- int
}

type Printer struct {
	in <-chan int
}

func main() {
	for {
		numbersChan := make(chan []int)
		averageChan := make(chan int)

		generator := Generator{out: numbersChan}
		calculator := Calculator{in: numbersChan, out: averageChan}
		printer := Printer{in: averageChan}

		go generator.generateRandomNumbers()
		go calculator.calculateAverage()
		go printer.printMean()

		time.Sleep(2 * time.Second)
	}
}

func (g Generator) generateRandomNumbers() {
	var randomNumbers []int

	for i := 0; i < 10; i++ {
		randomNumber := rand.Intn(100)
		randomNumbers = append(randomNumbers, randomNumber)
	}

	g.out <- randomNumbers

	close(g.out)
}

func (c Calculator) calculateAverage() {
	var sum int
	nums := <-c.in

	for _, n := range nums {
		sum += n
	}

	c.out <- (sum / len(nums))
	close(c.out)
}

func (p Printer) printMean() {
	fmt.Printf("Calculated mean value: %v\n", <-p.in)
}
