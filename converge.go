package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	canal := converge(trabalho("maçã"), trabalho("pêra"))
	for x := 0; x < 16; x++ {
		fmt.Println(<-canal)
	}

}

func trabalho(nome string) chan string {
	canal := make(chan string)
	go func(nome string, canal chan string) {
		for i := 1; ; i++ {
			canal <- fmt.Sprintf("Função %v diz: %v", nome, i)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)))
		}
	}(nome, canal)
	return canal
}

func converge(canal1, canal2 chan string) chan string {
	novoCanal := make(chan string)
	go func() {
		for {
			novoCanal <- <-canal1
		}
	}()
	go func() {
		for {
			novoCanal <- <-canal2
		}
	}()
	return novoCanal
}
