package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	fmt.Println("EXTERN PACKAGES")
	/*
		para importar paquete en consola escribir: go get github.com/google/uuid
		para importar version 1.22.0 del paquete en consola escribit: go get github.com/google/uuid@v1.22.0
		o cambiar version en archivo go.mod y actualizar en consola con: go mod tidy
	*/

	id1 := uuid.New()
	fmt.Println(id1)

	id2 := uuid.New().String()

	fmt.Println(id2)
}
