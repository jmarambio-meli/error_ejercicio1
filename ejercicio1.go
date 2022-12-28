package main

import (
	"fmt"
)

type myError struct {
	msg string
}

func (e *myError) Error() string {
	return "Error: El salario ingresado no alcanza el minimo Imponible"
}

func main() {
	e := myError{"Error"}
	var salary int = 160000
	if salary < 150000 {
		fmt.Println(e.Error())
		return
	}
	fmt.Println("Debe pagar impuesto")
}

func verSalario(salary int) (string, myError) {
	if salary < 150000 {
		return "", myError{"Error: El salario ingresado no alcanza el minimo Imponible"}
	}
	return "Debe pagar impuesto", myError{}
}
