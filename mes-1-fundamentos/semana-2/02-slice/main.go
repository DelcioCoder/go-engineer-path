package main

import (
	"fmt"
)

func main() {
	fmt.Println("Estrutura de dados Slice")
	//	Slice literal (mais comum)
	// nomes := []string{"Ana", "Dercio", "Bruno"}
	// fmt.Println(nomes)
	// fmt.Println(len(nomes))
	// fmt.Println(cap(nomes))

	// Usando make (muito usado quando se conhece o tamanho aproximado)
	// numeros := make([]int, 5)
	// fmt.Println(numeros)
	// fmt.Println(parte, len(parte), cap(parte))

	// Slice à partir de array ou outro Slice
	// tudo := []int{0, 10, 20, 30, 40, 50}
	// parte := tudo[1:5]
	// for i, v := range parte {
	// 	fmt.Println(i, v)
	// }

	// Regra final: O índice final é exclusivo
	// s := []int{0, 1, 2, 3, 4, 5}
	// fmt.Println(s[1:3]) //[1,2]
	// fmt.Println(s[:3])  //[0,1,2]
	// fmt.Println(s[3:])  //[3,4,5]
	// fmt.Println(s[2:2])

	// s := make([]int, 0, 5)
	// s = append(s, 1)
	// s = append(s, 2, 3)
	// s = append(s, 4, 5, 6)
	// fmt.Println(s, len(s), cap(s))

	// Armadilha clássica: Slice compartilhando o mesmo array
	// original := []int{1, 2, 3, 4, 5}
	// a := original[1:4]
	// b := original[2:5]
	// a[0] = 99
	// fmt.Println(original)
	// fmt.Println(b)

	// Usando o pacote slices Go 1.21+
    // s := []int{3, 1, 4, 1, 5, 9}
    // isContain := slices.Contains(s, 4)
    //    fmt.Println(isContain)
	// slices.Sort(s)
	// // Clone (cópia profunda)
	// copia := slices.Clone(s)
	// fmt.Println(len(copia))
	// fmt.Println(contem)
	// fmt.Println(s)
	// fmt.Println(copia)

	// copia := make([]int, len(s))
	// copy(copia, s)
	// fmt.Println("Slice copiado", copia)
}
