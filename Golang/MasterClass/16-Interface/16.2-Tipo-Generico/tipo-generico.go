package main

import (
	"fmt"
	"reflect"
)

func generica(interf interface{}) {
	fmt.Println(interf, " = ", reflect.TypeOf(interf))
}

func main() {
	generica(1)
	generica(0.4)
	generica("hahaha")

	saladaDaPorra := map[interface{}]interface{}{
		1:        "Carlos",
		"Manoel": true,
		false:    "Joaquina",
		3.1415:   "Pi",
	}
	fmt.Println(saladaDaPorra)
}
