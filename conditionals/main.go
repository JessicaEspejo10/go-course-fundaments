package main

import (
	"fmt"
)

func main() {
	yearsOld := 16

	fmt.Println("IF")

	if yearsOld > 18 {
		fmt.Println("Mayor a 18")
	} else {
		fmt.Println("Menor o igual a 18")
	}

	boolVar := true

	if boolVar {
		fmt.Println("Es verdadero")
	} else {
		fmt.Println("Es falso")
	}

	boolVar2 := false

	if boolVar2 {
		fmt.Println("Es verdadero")
	} else {
		fmt.Println("Es falso")
	}

	if value := true; value {
		fmt.Println("Es verdadero")
	}

	number := 10

	if number == 1 {
		fmt.Println("uno")
	} else if number == 2 {
		fmt.Println("dos")
	} else if number == 3 {
		fmt.Println("tres")
	} else {
		fmt.Println("ninguna es valida")
	}

	fmt.Println("SWITCH")

	switch number {
	case 1:
		fmt.Println("uno")
	case 2:
		fmt.Println("dos")
	case 3:
		fmt.Println("tres")
	case 4:
		fmt.Println("cuatro")
	default: //si ninguna se cumple
		fmt.Println("ninguna es válida")
	}

	switch {
	case number > 0:
		fmt.Println("positivo")
	case number < 0:
		fmt.Println("negativo")
	case number == 0:
		fmt.Println("cero")
	}

}
