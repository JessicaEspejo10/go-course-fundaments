package main

import "fmt"

func main() {
	//integer-string
	map1 := make(map[int]string)
	map1[1] = "a"
	map1[2] = "z"
	map1[7] = "e"
	map1[-10] = "jk"
	fmt.Println("map cpmpleto: ", map1)
	fmt.Println("key 7 del map: ", map1[7])

	//integer-array
	map2 := make(map[int][]string)
	map2[1] = []string{"z", "e"}
	map2[2] = []string{"z", "a"}
	map2[7] = []string{"e", "q"}
	map2[-10] = []string{"jk", "0"}
	fmt.Println("map2 cpmpleto: ", map2)
	fmt.Println("key 7 del map2: ", map2[7])

	//eliminar un elemento
	delete(map2, 7)
	fmt.Println("map cpmpleto: ", map2)

	//acceder a una key
	fmt.Println("key 2 del map2: ", map2[2])
	fmt.Println("key inexistente en el map2: ", map2[5])

	//consultar si un valor existe
	v, ok := map2[2]
	fmt.Println("El valor es: ", v, " - existe: ", ok)

	v2, ok2 := map2[8]
	fmt.Println("El valor es: ", v2, " - existe: ", ok2)

	//crear map
	map3 := map[int]string{
		4: "A",
		1: "D",
		9: "P",
	}
	fmt.Println(map3)
}
