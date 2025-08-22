package main

import (
	"AutomatedTests/enderecos"
	"fmt"
)

func main() {

	address := "Rua pastor Rodolfo Beuttemuller"

	type1 := enderecos.AddressType(address)
	fmt.Println(type1)

}
