package operacoes

import(
	cont"ac3/contatos"
	"fmt"
)

func AdicionaContato(c cont.Contato, a *[5]cont.Contato) {
	for ind, contato := range a {
		if (contato == cont.Contato{}) {
			a[ind] = c
			break
		}
	}
}
func RetornaLista() [5]cont.Contato{
	return [5]cont.Contato{}
}

func ExcluiContato(a *[5]cont.Contato) {
	if (a[0] == cont.Contato{}) {
		fmt.Println("Lista de contatos vazia!")
	}

	for ind, contato := range a {
		if (contato == cont.Contato{}) {
			a[ind-1] = cont.Contato{}
		}
	}

}

func EditaEmail(i int, novoEmail string, a [5]cont.Contato) {
	a[i].AlterarEmail(novoEmail)
	
}

