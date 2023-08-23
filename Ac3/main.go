package main

import (
	"bufio"
	"fmt"
	"os"
	"ac3/contatos"
	 ope "ac3/operacoes"
)


func main() {
	var listaContatos [5]contatos.Contato
	var nome, email, opcao, novoEmail string
	var i int

	leitor := bufio.NewReader(os.Stdin)

	fmt.Println("Lista de contatos!")
	for {
		fmt.Print("Digite (1) para adicionar, (2) para remover, (3) para exibir os contatos salvos, (4) para editar o email ou qualquer outra coisa para sair: ")
		fmt.Scanln(&opcao)

		switch opcao {
		case "1":
			fmt.Print("Informe o seu nome: ")
			nome, _ = leitor.ReadString('\n')

			fmt.Print("Informe o seu email: ")
			fmt.Scanln(&email)

			ope.AdicionaContato(contatos.Contato{Nome: nome, Email: email}, &listaContatos)
		case "2":
			ope.ExcluiContato(&listaContatos)
		case "3":
			ope.RetornaLista()
		case "4":
			ope.RetornaLista()
			fmt.Print("Informe o indice: ")
			fmt.Scanln(&i)

			fmt.Print("Informe o novo email: ")
			novoEmail, _ = leitor.ReadString('\n')

			ope.EditaEmail(i, novoEmail, listaContatos)
		default:
			return
		}

		fmt.Println(listaContatos)
	}

}




