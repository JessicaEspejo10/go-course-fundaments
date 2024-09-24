package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

func main() {

	fmt.Println("VARIABLES")

	//enteros
	var myIntVar int
	myIntVar = -12
	fmt.Println("My int variable is: ", myIntVar)

	//enterosPositivos
	var myUintVar uint
	myUintVar = 12
	fmt.Println("My uint variable is: ", myUintVar)

	//string
	var myStringVar string
	myStringVar = "Hello world"
	fmt.Println("My string variable is: ", myStringVar)

	//boolean
	var myBoolVar bool
	myBoolVar = true
	fmt.Println("My bool variable is: ", myBoolVar)

	//variables declaradas y no usadas generan error
	//var myOtherVar string

	//ver direccion de memorua de una variable
	fmt.Println("my variable address is: ", &myStringVar)

	//instanciar variable - declarar y asignar valor en una linea
	myIntVar2 := 12
	fmt.Println("my one line (:=) variable is: ", &myIntVar2)

	fmt.Println("CONSTANTS")

	const myIntConst int = 12
	fmt.Println("My int constant is: ", myIntConst)

	const myFirstIntConst = "a12"
	fmt.Println("My constant with no type is: ", myFirstIntConst)

	const myStringConst string = "test"
	fmt.Println("My string constant test: ", myStringConst)

	//BLOQUES
	fmt.Println("BLOQUES")

	myOuterScopeVar := 56
	{
		myScopeVar := 40
		fmt.Println("my scope var is: ", myScopeVar)
		fmt.Println("my outer scope var is: ", myOuterScopeVar)
	}

	//fmt.Println("my scope var is: ", myScopeVar)
	fmt.Println("my scope var cannot be read outside the scope")
	fmt.Println("my outer scope var is: ", myOuterScopeVar)

	fmt.Println("Uint VARIABLE SIZE")

	var my8BitsUnitVar uint8 = 90
	fmt.Printf("type: %T, value %d, bytes: %d, bits: %d \n", my8BitsUnitVar, my8BitsUnitVar, unsafe.Sizeof(my8BitsUnitVar), unsafe.Sizeof(my8BitsUnitVar)*8)

	var my16BitsUnitVar uint16 = 90
	fmt.Printf("type: %T, value %d, bytes: %d, bits: %d \n", my16BitsUnitVar, my16BitsUnitVar, unsafe.Sizeof(my16BitsUnitVar), unsafe.Sizeof(my16BitsUnitVar)*8)

	var my32BitsUnitVar uint32 = 90
	fmt.Printf("type: %T, value %d, bytes: %d, bits: %d \n", my32BitsUnitVar, my32BitsUnitVar, unsafe.Sizeof(my32BitsUnitVar), unsafe.Sizeof(my32BitsUnitVar)*8)

	var my64BitsUnitVar uint64 = 90
	fmt.Printf("type: %T, value %d, bytes: %d, bits: %d \n", my64BitsUnitVar, my64BitsUnitVar, unsafe.Sizeof(my64BitsUnitVar), unsafe.Sizeof(my64BitsUnitVar)*8)

	var myUintBitsUnitVar uint = 90
	fmt.Printf("type: %T, value %d, bytes: %d, bits: %d \n", myUintBitsUnitVar, myUintBitsUnitVar, unsafe.Sizeof(myUintBitsUnitVar), unsafe.Sizeof(myUintBitsUnitVar)*8)

	fmt.Println("FLOAT VARIABLE SIZE")

	var myFoat32Var float32 = 3.14
	fmt.Printf("type: %T, value: %f, bytes: %d, bits: %d \n", myFoat32Var, myFoat32Var, unsafe.Sizeof(myFoat32Var), unsafe.Sizeof(myFoat32Var)*8)

	var myFoat64Var float64 = 3.14
	fmt.Printf("type: %T, value: %f, bytes: %d, bits: %d \n", myFoat64Var, myFoat64Var, unsafe.Sizeof(myFoat64Var), unsafe.Sizeof(myFoat64Var)*8)

	fmt.Println("STRING VARIABLE SIZE")

	var myStringVar3 string = "test1234"
	fmt.Printf("value: %s \n", myStringVar3)

	myStringVar4 := `Mi varible es de tipo texto en go
	con multiples
	saltos
	de
	linea
	:)`
	fmt.Printf("value: %s \n", myStringVar4)

	fmt.Println("CONVERT TO STRING")

	{
		fmt.Println("float a string con Sprintf")
		floatVar := 33.25
		fmt.Printf("type: %T, value: %f \n", floatVar, floatVar)

		//convierte float a string
		floatStrVar := fmt.Sprintf("%.2f", floatVar)
		fmt.Printf("type: %T, value: %s \n", floatStrVar, floatStrVar)

		fmt.Println("int a string con Sprintf")

		intVar := 34
		fmt.Printf("type: %T, value: %d \n", intVar, intVar)

		//convierte int a string
		intStrVar := fmt.Sprintf("%d", intVar)
		fmt.Printf("type: %T, value: %s \n", intStrVar, intStrVar)

		fmt.Println("int a string con strconv package")

		intVar2 := 342
		intStrVar2 := strconv.Itoa(intVar2)
		fmt.Printf("type: %T, value: %s \n", intStrVar2, intStrVar2)
	}

	fmt.Println("CONVERT TO NUMBER")

	{
		strIntVar2, err := strconv.Atoi("15")
		fmt.Printf("type: %T, value: %d, err: %v \n", strIntVar2, strIntVar2, err)

		strIntVar3, err := strconv.Atoi("15Td")
		fmt.Printf("type: %T, value: %d, err: %v \n", strIntVar3, strIntVar3, err)

		//segundo parametro es la base (decimal, hexadecimal, ....) y el tercero el tamaño en bits
		//float se pueden pasar a String con strconv.ParseFloat
		strIntVar4, err := strconv.ParseInt("65", 10, 64)
		fmt.Printf("type: %T, value: %d, err: %v \n", strIntVar4, strIntVar4, err)

	}
}
