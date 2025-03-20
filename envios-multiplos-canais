package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	canal := make(chan int)

	for range 10 {
		envia(canal, 10)

	}

	go recebe(canal)
	wg.Wait()
	close(canal)

}

func envia(canal chan<- int, qtd int) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range qtd {
			canal <- v
		}
	}()
}

func recebe(canal <-chan int) {
	total := 1
	for range canal {
		fmt.Println(total)
		total++
	}
}
