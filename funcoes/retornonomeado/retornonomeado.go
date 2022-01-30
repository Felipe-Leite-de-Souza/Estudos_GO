package main

import "fmt"

func trocar(p1, p2 int) (segundo int, primeiro int) {
	segundo = p2
	primeiro = p1
	return //retorno limpo
}

func main() {

	//podemos usar quantas vezes quisermos
	x, y := trocar(2, 3)
	fmt.Println(x, y)

	//usar uma úca vez por conta do nome das variáveis
	segundo, primeiro := trocar(7, 1)
	fmt.Println(segundo, primeiro)
}