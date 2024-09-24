package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("BUCLES")
	sum := 0
	for i := 0; i < 10; i++ {
		//fmt.Println("i es " + strconv.Itoa(i))
		sum++
	}
	fmt.Println("El total es " + strconv.Itoa(sum))

	sum = 0
	fmt.Println("Al reiniciar la suma el valor es " + strconv.Itoa(sum))

	for sum < 1000 {
		sum++
	}
	fmt.Println("El total es " + strconv.Itoa(sum))

	sum = 0

	for {
		if sum > 1000 {
			break
		}
		sum++
	}
	fmt.Println("El total es " + strconv.Itoa(sum))

}
