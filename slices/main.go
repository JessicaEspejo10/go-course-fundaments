package main

import (
	"fmt"
)

func main() {
	fmt.Println("SLICES")
	myArrayVar := [5]int{3, 6, 9, 12, 15}
	fmt.Println("array: ", myArrayVar, " - len: ", len(myArrayVar))

	//definir slice
	mySliceVar := []int{}
	fmt.Println("slice: ", mySliceVar, " - len: ", len(mySliceVar))

	//añadir valores
	mySliceVar = append(mySliceVar, 18, 21, 44, 89, 24)
	fmt.Println("slice: ", mySliceVar, " - len: ", len(mySliceVar))

	//tomar una parte de otro array o slice
	mySliceVar2 := []int{}
	mySliceVar2 = myArrayVar[2:4]
	fmt.Println("slice2: ", mySliceVar2, " - len: ", len(mySliceVar2))

	//modificar el slice modifica la posicion del array que se tomó para crearlo
	mySliceVar2[0] = 58
	mySliceVar2[1] = 49
	fmt.Println("slice2: ", mySliceVar2, " - len: ", len(mySliceVar2), " - address: ", &mySliceVar2[0])
	fmt.Println("array: ", myArrayVar, " - len: ", len(myArrayVar), " - address: ", &myArrayVar[0])

	mySliceVar3 := mySliceVar[:3]
	fmt.Println("slice3 de posición 0 a 3: ", mySliceVar3)
	mySliceVar4 := mySliceVar[2:]
	fmt.Println("slice4 de posición 2 a final: ", mySliceVar4)

	//inicializar de otra forma
	mySliceVar5 := make([]int, 7)
	fmt.Println("slice5: ", mySliceVar5, " - len: ", len(mySliceVar5), " - address: ", &mySliceVar5[0])

	mySliceVar6 := []string{"a", "c", "e", "g", "i", "k"}
	fmt.Println("slice6: ", mySliceVar6, " - len: ", len(mySliceVar6), " - address: ", &mySliceVar6[0])

}
