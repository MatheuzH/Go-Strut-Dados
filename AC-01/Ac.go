package main

import "fmt"

func main() {
	fmt.Println(semana(3))
}


func semana(num int) string{
	switch num{
	case 1:
		return "Domingo!"
	case 2:
		return "Segunda!"
	case 3:
		return "Terça!"
	case 4:
		return "Quarta!"
	case 5:
		return "Quinta!"
	case 6:
		return "Sexta!"
	case 7:
		return "Sabado!"
	default:
		return "Valor inválido!"
	}
}

func e_bissextp(ano int) bool {
	bissexto := false
	for i := 4; i %  {

	}
	return bissexto
}
