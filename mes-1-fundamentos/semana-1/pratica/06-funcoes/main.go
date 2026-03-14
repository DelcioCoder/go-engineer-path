package main

import "fmt"

// func sum(a, b int) int {
// 	return a + b
// }
// result := sum(10, 20)
// fmt.Println(result)

// func division(a, b float64) (float64, error) {
// 	if b == 2 {
// 		return 0, fmt.Errorf("Divisão por zero")
// 	}
//
// resultado, err := division(10, 2)
// if err != nil {
// 	fmt.Println(err)
// 	return
// }
//
// }

func main() {
	somar := func(a, b int) int {
		return a + b
	}

	fmt.Println(somar(10, 2))
}
