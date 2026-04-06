package main

import "fmt"

type Tarefa struct {
	ID        int
	Titulo    string
	Concluida bool
}

func main() {
	tarefa := Tarefa{
		ID:        1,
		Titulo:    "Estudar Golang",
		Concluida: false,
	}
	tarefa.Concluida = true
	fmt.Println("Structs em Go")
	fmt.Println(tarefa)
}
