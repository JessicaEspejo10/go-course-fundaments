package main

import "fmt"

type MyStruct struct {
	ID   int
	Name string
}

func (m MyStruct) Set(name string) {
	fmt.Printf("address: %p \n", &m)
	m.Name = name
}

func (m *MyStruct) SetP(name string) {
	fmt.Printf("address: %p \n", m)
	m.Name = name
}

func main() {
	var i int
	var iP *int
	i = 34

	iP = &i

	fmt.Println("var value : ", i, " - address: ", &i)
	fmt.Println("pointer value: ", iP, " - address: ", &iP)

	*iP = 1
	//como el puntero apunta a la dirección de variable, al modificarlo se modifica la variable original a la que apunta
	fmt.Printf("value i: %d, value pointer: %d \n", i, *iP)

	myVar := 30
	fmt.Printf("var value : %d, address: %p \n", myVar, &myVar)

	fmt.Println("-----------------------------------")
	increment(myVar)
	fmt.Printf("var value : %d, address: %p \n", myVar, &myVar)

	fmt.Println("-----------------------------------")
	incrementPointer(&myVar)
	fmt.Printf("var value : %d, address: %p \n", myVar, &myVar)

	var mySlice []int
	mySlice = append(mySlice, 3, 4, 2)
	fmt.Printf("address: %p, values: %v \n", mySlice, mySlice)
	fmt.Printf("addr1: %p, val1: %v \n", &mySlice[0], mySlice[0])
	fmt.Printf("addr2: %p, val2: %v \n", &mySlice[1], mySlice[1])
	fmt.Printf("addr3: %p, val3: %v \n", &mySlice[2], mySlice[2])

	updateSlice(mySlice)
	fmt.Println(mySlice)

	myStruct := &MyStruct{ID: 123, Name: "Test"}
	fmt.Printf("address: %p \n", myStruct)
	fmt.Printf("id: %d, name %s \n", myStruct.ID, myStruct.Name)

	updateStruct(myStruct)
	fmt.Printf("id: %d, name %s \n", myStruct.ID, myStruct.Name)

	myStruct.Set("test method")
	fmt.Printf("id: %d, name %s \n", myStruct.ID, myStruct.Name)

	myStruct.SetP("test method 2")
	fmt.Printf("id: %d, name %s \n", myStruct.ID, myStruct.Name)
}

func increment(val int) {
	val++
	fmt.Println("var value : ", val, " - address: ", &val)
}

func incrementPointer(val *int) {
	*val++
	fmt.Println("pointer value : ", val, " - address: ", &val, " - pointer: ", *val)
}

func updateSlice(v []int) {
	fmt.Printf("address: %p \n", v)
	fmt.Printf("address1: %p, address2: %p, address3: %p \n ", &v[0], &v[1], &v[2])
	v[0] = 12
}

func updateStruct(v *MyStruct) {
	fmt.Printf("address in fuction: %p \n", v)
	v.ID = 998
	v.Name = "Margarita"

}
