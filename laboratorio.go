package main

import "fmt"

func main() {
	var numero int
	fmt.Println("ingrese un numero")
	fmt.Scan(&numero)
	fmt.Printf("El Exponencial(%d): %d ", numero, factorial(numero))
}

func factorial(num int) int {
	if num == 0 {
		return 1
	}
	resultado := num * factorial(num-1)
	return resultado
}
