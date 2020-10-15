package main

import "fmt"

func main() {

	var s []int
	var x int = 0
	var n int
	var a int
	fmt.Println("Cuantos numeros quiere agregar: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Println("Numero: ")
		fmt.Scan(&a)
		s = append(s, a)
		x = x + s[i]

	}
	fmt.Println("Resultado: ", x)

}
