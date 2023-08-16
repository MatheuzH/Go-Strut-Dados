package main

import "fmt"

type Contato struct{
	Nome	string
	Email	string
}

func (c *Contato) MudarEmail(novoEmail string) {
	c.Email = novoEmail
}

func AdicionaContato(c Contato, cont [5]Contato) [5]Contato {
	for i, c := range cont {
		if (c == Contato{}) {
			cont[i] = c
			break
		}
	}
	return cont
}

func ExcluiContato(cont [5]Contato) [5]Contato{
	for i, contato := range cont {
		if (contato == Contato{}){
			cont[i - 1] = Contato{}
			break
		}
	}
	return cont
}

func main() {
	var contato [5]Contato
	var x int
	var y Contato
	fmt.Println("Selecione uma opção")
	fmt.Println("1 - Criar Contato")
	fmt.Println("2 - Apagar Contato")
	fmt.Println("3 - Sair")
	fmt.Scanln(x)

	switch x {
	case 1:
		fmt.Println("Informe o nome e o email")
		fmt.Scanln(y)
		AdicionaContato(y, contato)
	}
}

//Não aguento mais, TENTEI!!!!