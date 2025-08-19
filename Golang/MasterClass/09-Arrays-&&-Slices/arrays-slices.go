package main

import (
	"fmt"
	"reflect"
)

func main() {

	var array1 [5]string
	array1[0] = "1st Position"
	fmt.Println(array1)

	array2 := [5]string{"1st Position", "2nd Position", "3rd Position", "4th Position", "5th Position"}
	fmt.Println(array2)

	array3 := [...]int16{1, 2, 3, 4, 5}
	fmt.Println(array3)

	// o ARRAY é uma estrutura mais rígida, não permite alteração de tamanho sobre demanda;
	// por isso em GO é mais utilizado o SLICE. Pois permite alterar o tamanho conforme desejado.

	slice := []int16{10, 11, 12, 13, 14, 15, 16, 17}
	fmt.Println(slice)

	// ARRAYs e SLICEs são tipos diferentes de dados;
	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	slice = append(slice, 9)
	fmt.Println(slice)

	// ao apontar um intervalo de indices "[1:3]" o primeiro é inclusive e o segundo é exclusivo
	// então no exemplo acima apenas os conteúdos dos indices 1 e 2 serão importados.
	slice2 := array2[1:3]
	fmt.Println(slice2)

	// o Slice é um ponteiro para indices de um vetor.
	array2[1] = "2a Posição"
	fmt.Println(slice2)

	arrayInterno()
}

func arrayInterno() {
	// ARRAY Interno
	fmt.Println("--------")
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)
	fmt.Println(len(slice3), cap(slice3))
	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6)
	fmt.Println(len(slice3), cap(slice3))

	// quando o limite do slice atingido, é criado um novo array com o dobro do tamanho atual
	// QUANDO ATINGE A META, ELE DOBRA A META.
}
