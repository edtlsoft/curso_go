package main

import (
	"fmt"
	"github.com/edtlsoft/GoCalculator"
)

func main(){
	number1  := GoCalculator.LeerEntrada("Ingrese el primer numero")
	operador := GoCalculator.LeerEntrada("Ingrese la operación a realizar ('+', '-', '*', '/')")
	number2  := GoCalculator.LeerEntrada("Ingrese el segundo numero")

	calculator 	:= GoCalculator.Calculator{} 
	result 		:= calculator.Operate(number1, number2, operador)

	fmt.Println(number1, operador, number2)
	fmt.Println(result)
}