package main

import (
	"fmt"
)

func main() {
	n := []int{1, 2, 3, 4, 5}
	alvo := 7

	num1, num2 := ArraySoma(n, alvo)
	if num1 != -1 && num2 != -1 {
		fmt.Println("Par encontrado:", num1, "e", num2)
	} else {
		fmt.Println("-1 -1, nenhum par foi encontrado")
	}
}

func ArraySoma(n []int, alvo int) (int, int) {
	inicio := 0
	final := len(n) - 1

	for inicio < final {
		soma := n[inicio] + n[final]
		if soma == alvo {
			return n[inicio], n[final]
		} else if soma < alvo {
			inicio++
		} else {
			final--
		}
	}

	return -1, -1
}
