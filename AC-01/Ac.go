package main

import "fmt"

func main() {
	fmt.Println(semana(3))
	fmt.Println(e_bissexto(2000))

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

func e_bissexto(num int) string{
	if num % 4 == 0{
		if num % 100 == 0{
			if num % 400 == 0 {
				return "É bissexto"
			} else {
				return "Não é bissexto"
			}
		} else{
			return "É bissexto"
		} 
	} else {
		return "Não é bissexto"
	}

}


