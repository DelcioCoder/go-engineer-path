package main

import "fmt"

func main() {
	fmt.Println("Aprender a declarar variaveis em Go")
	var (
		contador int
		nome string
		lista []int
	)

	// Forma explicita de declarar variáveis em Go
	var name string = "Dercio"
	fmt.Println("Nome: ", name)


	// Declaração com inferência de tipo 
	var idade = 22
	fmt.Println("Idade: ", idade)

	// Declaração curta(short declaration)
	cidade := "Luanda"
	activo := true
	fmt.Println("Cidade: ", cidade)
	fmt.Println("Activo: ", activo)

	// Zero Values
	fmt.Println(contador)
	fmt.Println(nome)
	fmt.Println(lista)
}
