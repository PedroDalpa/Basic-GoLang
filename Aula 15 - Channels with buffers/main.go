package main

import (
	"fmt"
	"time"
)

func numbers(n chan<- int) {
	for i := 0; i < 10; i++ {
		n <- i
		fmt.Printf("Escrito no channel: %d \n", i)
		time.Sleep(time.Millisecond * 50)
	}
	close(n)
}

func main() {
	cn := make(chan int, 3)

	go numbers(cn)

	for v := range cn {
		fmt.Printf("Lido do channel: %d \n", v)
		time.Sleep(time.Millisecond * 150)
	}

	fmt.Println("Fim da execucao")
}
