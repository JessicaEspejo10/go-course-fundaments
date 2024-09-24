package main

import "fmt"

func main() {
	v1 := []float64{1.3, 5.45, 12.223, 6.92, 78.102}
	v2 := []int32{9, 23, 1, 23, 8, 98}
	fmt.Println(Sum01(v1))
	fmt.Println(Sum01(v2))

	fmt.Println(Sum02(v1))
	fmt.Println(Sum02(v2))

	//el tipo de datos genericos de los dos argumentos debe ser igual
	anyType("1", "1.2")

	//esta funcion si recibe dos genericos de diferente tipo, de acuerdo con los argumentos de la función
	anyTypeTwo("3", 1.025)

}

//se especifican los tipos de datos a recibir. en este caso recibe un array generico que recibe loso tipos que estan entre [] y devuelve un unico valor generico
func Sum01[N int | int32 | int64 | float32 | float64](v []N) N {
	var total N
	for _, tV := range v {
		total += tV
	}
	return total
}

type Number interface {
	int | int32 | int64 | float32 | float64
}

func Sum02[N Number](v []N) N {
	var total N
	for _, tV := range v {
		total += tV
	}
	return total
}

func anyType[N any](v1, v2 N) {
	fmt.Println(v1, v2)
}

func anyTypeTwo[N1 any, N2 any](v1 N1, v2 N2) {
	fmt.Println(v1, v2)
}
