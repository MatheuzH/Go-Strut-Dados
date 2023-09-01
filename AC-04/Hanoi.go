package main

import (
	"fmt"
)

func main() {
	numDiscos := 2
	movimentos := hanoi(numDiscos, "Origem", "Trabalho", "Destino")
	fmt.Println("Foram necessários", movimentos, "movimentos")
}

func hanoi(n int, origem, trabalho, destino string) int {
	if n != 0 {
		movimentos := hanoi(n-1, origem, destino, trabalho)

		fmt.Println("Mova o disco", n, "de", origem, "para", destino)
		movimentos++
		movimentos += hanoi(n-1, trabalho, origem, destino)
		return movimentos
	} else {
		return 0
	}
}
