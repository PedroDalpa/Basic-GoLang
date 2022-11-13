package main

import (
	"fmt"
	"time"
)

func numbers() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d \n", i)
		time.Sleep(time.Millisecond * 150)
	}
}

func letters() {
	for i := 'a'; i < 'j'; i++ {
		fmt.Printf("%c \n", i)
		time.Sleep(time.Millisecond * 230)
	}
}

func main() {
	go numbers()
	go letters()
	time.Sleep(5 * time.Second)
	fmt.Println("Fim da execucao")
}
