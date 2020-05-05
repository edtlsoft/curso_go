package main

import (
	"fmt"
)

func cambiarValor(valor *int) {
	fmt.Println("cambiarValor: ", valor, *valor)
	*valor = 50
}

func main() {
	x := 25

	fmt.Println(x)
	fmt.Println(&x)
	fmt.Println(*&x)
	
	cambiarValor(&x)
	fmt.Println("cambiarValor(x): ", x)

	y := &x
	fmt.Println(y)
	fmt.Println(*y)
}
