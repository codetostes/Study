package main

import "fmt"

func main() {
	tarefas, resultados := make(chan int, 45), make(chan int, 45)

	workers := 1
	for i := 0; i < workers; i++ {
		go worker(tarefas, resultados)
	}

	for i := 0; i < 45; i++ {
		tarefas <- i
	}
	close(tarefas)

	for i := 0; i < 45; i++ {
		fmt.Println(<-resultados)
	}
}

func worker(tarefas <-chan int, resultado chan<- int) {

	for t := range tarefas {
		resultado <- fibonacci(t)
	}

}

func fibonacci(pos int) int {
	if pos <= 1 {
		return pos
	}
	return fibonacci(pos-1) + fibonacci(pos-2)
}
