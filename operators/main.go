package main

import (
	"fmt"
)

func main() {

	fmt.Println("OPERADORES")

	yearsOld := 63

	fmt.Println(yearsOld > 30)  //true
	fmt.Println(yearsOld < 30)  //false
	fmt.Println(yearsOld <= 63) //true
	fmt.Println(yearsOld == 33) //false

	fmt.Println("OR")                            //true
	fmt.Println(yearsOld < 33 || yearsOld == 63) //F + T = true
	fmt.Println(yearsOld < 33 || yearsOld == 34) // F + F = false
	fmt.Println(yearsOld > 33 || yearsOld == 63) // T + T = true

	fmt.Println("AND")                           //true
	fmt.Println(yearsOld < 33 && yearsOld == 63) //F + T = false
	fmt.Println(yearsOld < 33 && yearsOld == 34) // F + F = false
	fmt.Println(yearsOld > 33 && yearsOld == 63) // T + T = true

	fmt.Println("NOT")                              //true
	fmt.Println(!(yearsOld < 33 && yearsOld == 63)) // NOT(F + T) = NOT(false) = true
	fmt.Println(!(yearsOld < 33 && yearsOld == 34)) // NOT(F + F) = NOT(false) = true
	fmt.Println(!(yearsOld > 33 && yearsOld == 63)) // NOT(T + T) = NOT(true) = false

	fmt.Println("COMBINACIONES")                                    //true
	fmt.Println(yearsOld < 33 && yearsOld == 63 || yearsOld > 40)   //F AND T OR T = true
	fmt.Println(yearsOld < 33 && (yearsOld == 63 || yearsOld > 40)) //F AND (T OR T) = false
}
