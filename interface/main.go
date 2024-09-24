package main

import (
	"fmt"
	"fundaments/interface/vehicles"
)

func main() {
	vArray := []string{"CAR", "MOTORCYCLE", "TRUCK", "MOTORCYCLE", "PLANE"}

	var d float64
	for _, v := range vArray {
		fmt.Printf("Vechicle: %s \n", v)
		veh, err := vehicles.New(v, 400)

		if err != nil {
			fmt.Println("Error: ", err)
			fmt.Println()
			continue
		}

		distance := veh.Distance()
		fmt.Printf("Distance: %.2f \n", distance)
		fmt.Println()
		d += distance
	}
	fmt.Printf("Total distance: %.2f \n", d)
}

func Display(value interface{}) {
	fmt.Println(value)
}
