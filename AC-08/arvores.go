package main

import (
	"fmt"
)

type No struct {
	Raiz int
	Esq  *No
	Dir  *No
}

type ArvoreBinaria struct {
	Arvore *No
}


func inserirNo(arvore *No, chave int) *No {
	if arvore == nil {
		return &No{Raiz: chave, Esq: nil, Dir: nil}
	}

	if chave < arvore.Raiz {
		arvore.Esq = inserirNo(arvore.Esq, chave)
	} else if chave > arvore.Raiz {
		arvore.Dir = inserirNo(arvore.Dir, chave)
	}

	return arvore
}

func (arvore *ArvoreBinaria) Inserir(chave int) {
	arvore.Arvore = inserirNo(arvore.Arvore, chave)
}

func buscaRaiz(arvore *No, chave int) bool {
	if arvore == nil {
		return false
	}

	if arvore.Raiz == chave {
		return true
	} else if chave < arvore.Raiz {
		return buscaRaiz(arvore.Esq, chave)
	} else {
		return buscaRaiz(arvore.Dir, chave)
	}
}

func (arvore *ArvoreBinaria) Busca(chave int) bool {
	return buscaRaiz(arvore.Arvore, chave)
}


func (arvore *ArvoreBinaria) OrdemCrescente() {
	ordem(arvore.Arvore)
}

func ordem(arvore *No) {
	if arvore != nil {
		ordem(arvore.Esq)
		fmt.Println(arvore.Raiz)
		ordem(arvore.Dir)
	}
}

func main() {
	arvoreBinaria := ArvoreBinaria{}

	// Inserção de elementos na árvore
	arvoreBinaria.Inserir(50)
	arvoreBinaria.Inserir(30)
	arvoreBinaria.Inserir(20)
	arvoreBinaria.Inserir(40)
	arvoreBinaria.Inserir(70)
	arvoreBinaria.Inserir(60)
	arvoreBinaria.Inserir(80)
	// BirInserirsca por um valor na árvore
	chave := 60
	if arvoreBinaria.Busca(chave) {
		fmt.Printf("Chave %d encontrada na árvore.\n", chave)
	} else {
		fmt.Printf("Chave %d não encontrada na árvore.\n", chave)
	}

	// Imprimir a árvore em ordem
	fmt.Println("Árvore em ordem:")
	arvoreBinaria.OrdemCrescente()
}
