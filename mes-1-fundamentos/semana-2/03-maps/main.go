package main

import (
	"fmt"
)

func Average(sum float64, a float64) float64 {
	return sum / a
}

func MaxGrade(studentGrade map[string]float64) float64 {
	maxGrade := studentGrade["student_1"]
	for _, v := range studentGrade {
		if v > maxGrade {
			maxGrade = v
		}
	}
	return maxGrade
}

func MinGrade(studentGrade map[string]float64) float64 {
	minGrade := studentGrade["student_1"]
	for _, v := range studentGrade {
		if v < minGrade {
			minGrade = v
		}
	}
	return minGrade
}

func main() {
	fmt.Println("Estrutura de dados Maps")

	// Literals maps
	// users := map[string]string{
	// 	"user_1": "Dercio",
	// 	"user_2": "Armando",
	// }
	// fmt.Println(users)
	// for k, v := range users {
	// 	fmt.Println(k, v)
	//
	// }

	// Forma idiomática de declarar maps(a mais usada)
	// usando o `make`
	// users := make(map[string]string)

	// Adicionar valores
	// users["user_1"] = "Dercio"
	// users["user_2"] = "Armando"
	// Remover elemento
	// delete(users, "user_2")
	// fmt.Println("Maps usando make")
	// fmt.Println(users)
	// Acessar valores
	// derciosName := users["user_1"]
	// fmt.Printf("O nome na posição user_1 é %s\n ", derciosName)

	// Verificar se o nome existe
	// name, ok := users["user_1"]
	// if ok {
	// 	fmt.Println("Nome encontrado:", name)
	// } else {
	// 	fmt.Println("Nome não encontrado!!")
	// }

	// Forma mais idiomática de verificar
	// if name, ok := users["user_1"]; ok {
	// 	fmt.Println("Encontrado: ", name)
	// }

	// Reference type
	// a := map[string]int{"x": 1}
	// fmt.Println(a)
	// b := a
	// b["x"] = 100
	// fmt.Println(a)

	// Desafio: Criar um programa que armazena notas de alunos em maps e calcule a média, maior
	// e a menor nota

	// var average float64
	var sum float64

	studentGrade := map[string]float64{
		"student_1": 10.5,
		"student_2": 11.5,
		"student_3": 20,
	}
	a := float64(len(studentGrade))

	for _, v := range studentGrade {
		sum += v
	}
	fmt.Println("A média das notas é:", Average(sum, a))
	fmt.Println("A nota máxima é:", MaxGrade(studentGrade))
	fmt.Println("A nota mínima é:", MinGrade(studentGrade))

}
