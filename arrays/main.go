package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("ARRAYS")

	var myIntVar int
	fmt.Printf("type: %T, value: %d, bytes: %d, bits: %d \n", myIntVar, myIntVar, unsafe.Sizeof(myIntVar), unsafe.Sizeof(myIntVar)*8)

	var myArrayVar1 [5]int //genera array de 5 posiciones
	fmt.Println(myArrayVar1, " - Size: ", len(myArrayVar1))

	myArrayVar2 := [3]string{"a", "b", "c"} //genera array de 3 posiciones definidas
	fmt.Println(myArrayVar2, " - Size: ", len(myArrayVar2))

	//asignar valores a posiciones de array
	myArrayVar1[0] = 1
	myArrayVar1[2] = 3
	myArrayVar1[4] = 5
	fmt.Println(myArrayVar1, " - Size: ", len(myArrayVar1))
	fmt.Printf("type: %T, value: %d, bytes: %d, bits: %d \n", myArrayVar1, myArrayVar1, unsafe.Sizeof(myArrayVar1), unsafe.Sizeof(myArrayVar1)*8)

	//recorrer array

	for i, v := range myArrayVar1 {
		fmt.Println("index: ", i, " - value: ", v)
	}
}
