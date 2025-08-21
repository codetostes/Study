package main

import (
	"fmt"
	"tdd/enderecos"
)

func main() {

	address := "Rua pastor Rodolfo Beuttemuller"

	type1 := enderecos.AddressType(address)
	fmt.Println(type1)

}
