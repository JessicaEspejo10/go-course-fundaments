package main

import (
	"fmt"
	"fundaments/functions/function"
)

func main() {
	var myIntVar int64 = 35
	function.Display(myIntVar)
	var myIntVar2 int64 = 56
	function.Display(myIntVar2)
	var myIntVar3 int64 = 78
	function.Display(myIntVar3)

	var a int = 64
	var b int = 48
	fmt.Println("La suma da ", function.Add(a, b))
	fmt.Println()

	function.RepeatString(10, "LA ")

	value, err := function.Calc(function.SUM, 20.12, 58)
	fmt.Println("value: ", value, " - error: ", err)

	value2, err2 := function.Calc(function.SUB, 20.12, 58)
	fmt.Println("value: ", value2, " - error: ", err2)

	value3, err3 := function.Calc(function.MUL, 20.12, 58)
	fmt.Println("value: ", value3, " - error: ", err3)

	value4, err4 := function.Calc(function.DIV, 20.12, 0)
	fmt.Println("value: ", value4, " - error: ", err4)

	value5, err5 := function.Calc(function.TEST, 20.12, 0)
	fmt.Println("value: ", value5, " - error: ", err5)

	xVal, yVal := function.Split(50)
	fmt.Println(xVal, " y ", yVal)

	val2 := function.Msum(1, 5, 8, 6, 9, 5, 3)
	fmt.Println(val2)

	val3 := function.Msum(1, 10, 27)
	fmt.Println(val3)

	mOpValue, err := function.MOperations(function.MUL, 4, 5, 7, 8, 9, 1)
	fmt.Println("value: ", mOpValue, " - err: ", err)

	factoryFunction := function.FactoryOperation(function.MUL)
	fmt.Println(factoryFunction(10,5))
}
