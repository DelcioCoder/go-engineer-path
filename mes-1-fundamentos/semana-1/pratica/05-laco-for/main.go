package main

import "fmt"

func main() {
	fmt.Println("Estudo do laço for único")
	// for i := 0; i < 100; i++ {
	// 	fmt.Println(i)
	// }

	// // Loop estilo while em outras linguagens
	// for i < 100 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// i := 0
	// var name string
	// var opc string
	//
	// for {
	// 	fmt.Println("Digitem os vossos nomes: ")
	// 	fmt.Scan(&name)
	// 	fmt.Println("Deseja parar? (s/n)")
	// 	fmt.Scan(&opc)
	// 	i++
	//
	// 	if opc == "s" {
	// 		continue
	// 	} else if opc == "n" {
	// 		break
	// 	} else {
	// 		fmt.Println("Opção inválida!!")
	// 	}
	// }

	users := []string{"Dercio", "Armando"}

	for i, v := range users {
		fmt.Println(i, v)
	}

}
