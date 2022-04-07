package main

import (
	"fmt"
)

func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
	}

	close(out)
}

func square(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}

	close(out)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	na := make(chan int)
	sq := make(chan int)

	go counter(na)
	go square(sq, na)
	printer(sq)
	
}


func mirroredQuery() string {
	response := make(chan string, 3)

	go func() { response <- request("1") }()
	go func() { response <- request("2") }()
	go func() { response <- request("3") }()

	return <- response
}

func request(hostname string) (response string) {
	
	return ""
}