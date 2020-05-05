package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

type calc struct {}

func (calc) operate(number1 string, number2 string, operador string) int {
	response := 0

	operador1 := parsear(number1) 
	operador2 := parsear(number2) 

	switch operador {
		case "+":
			response = operador1 + operador2
			break
		case "-":
			response = operador1 - operador2
			break
		case "*":
			response = operador1 * operador2
			break
		case "/":
			response = operador1 / operador2
			break
		default:
			response = 0
	}

	return response
}

func parsear (entrada string) int {
	number, error := strconv.Atoi(entrada)

	if error != nil {
		fmt.Printf("Ocurrio un error al intentar parsear la entrada: %s, valor de salida: %d", error, number)
	}

	return number
}

func leerEntrada(mensaje string) string {
	fmt.Println(mensaje)
	
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	entrada := scanner.Text()

	return entrada
}

func main() {
	fmt.Println("C A L C U L A D O R A  -  G O")

	number1  := leerEntrada("Ingrese el primer numero")
	operador := leerEntrada("Ingrese la operación a realizar ('+', '-', '*', '/')")
	number2  := leerEntrada("Ingrese el segundo numero")

	c := calc{} 
	result := c.operate(number1, number2, operador)

	fmt.Println(number1, operador, number2)

	fmt.Println(result)
}