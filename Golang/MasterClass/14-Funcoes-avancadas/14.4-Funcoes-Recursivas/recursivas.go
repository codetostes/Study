package main

import "fmt"

func fibonacci(pos uint) uint {
	if pos <= 1 {
		return pos
	}
	return fibonacci(pos-1) + fibonacci(pos-2)

}
func main() {

	posicao := uint(10)
	for i := uint(0); i <= posicao; i++ {
		fmt.Printf("posição %d do bonacci = %d\n", i, fibonacci(i))
	}
}
