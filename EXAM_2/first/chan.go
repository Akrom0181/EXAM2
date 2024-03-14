package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sendNum(channel chan int) {
	min, max := 10, 30
	minSend, maxSend := 5, 10
	random := rand.Intn(maxSend-minSend) + minSend
	fmt.Printf("Sending %d random numbers\n", random)
	for i := 0; i < random; i++ {
		number := rand.Intn(max-min) + min
		fmt.Println("Sending number:", number)
		channel <- number
		time.Sleep(time.Millisecond * 700)
	}
	close(channel)
}

func squareNumbers(input <-chan int, output chan<- int) {
	for num := range input {
		squared := num * num
		output <- squared
	}
	close(output)
}

func main() {

	channel1 := make(chan int)
	channel2 := make(chan int)

	go sendNum(channel1)
	go squareNumbers(channel1, channel2)

	for squaredNum := range channel2 {
		fmt.Println("Squared number:", squaredNum)
	}

}

