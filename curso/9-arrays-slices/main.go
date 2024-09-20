package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")
	fmt.Println("-----------------")

	var array1 [5]int
	array1[0] = 1

	fmt.Println(array1)

	fmt.Println("-----------------")

	array2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array2)

	fmt.Println("-----------------")

	slice1 := []int{1, 2, 3, 4, 5} // slice é um array dinâmico
	fmt.Println(slice1)

	fmt.Println("-----------------")

	fmt.Println(slice1[1:3]) // fatiando um slice
	fmt.Println(slice1[:3])  // fatiando um slice

	fmt.Println("-----------------")
	fmt.Println(reflect.TypeOf(slice1))
	fmt.Println(reflect.TypeOf(array2))

	fmt.Println("-----------------")
	slice1 = append(slice1, 6)
	fmt.Println(slice1)

}
