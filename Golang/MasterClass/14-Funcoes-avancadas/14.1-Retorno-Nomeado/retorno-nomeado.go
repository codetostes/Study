package main

import "fmt"

func mathCalcs(n1, n2 int) (sum, sub int) {
	sum = n1 + n2
	sub = n1 - n2
	return
}

func main() {
	fmt.Println(mathCalcs(10, 5))

}
