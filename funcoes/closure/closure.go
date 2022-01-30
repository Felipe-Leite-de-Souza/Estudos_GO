package main

import "fmt"

func closure() func() {
	x := 10
	var funcacao = func() {
		fmt.Println(x)
	}
	return funcacao
}

func main() {
	x := 20
	fmt.Println(x)

	imprimeX := closure()
	imprimeX()
}
