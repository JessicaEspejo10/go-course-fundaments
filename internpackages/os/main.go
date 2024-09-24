package main

import (
	"fmt"
	"os" //usa funcionalidades del sistema operativo
)

func main() {
	file, err := os.Open("main.go")
	if err != nil {
		fmt.Println(err)
		os.Exit(1) //0 exitoso, 1 fallido
	}
	fmt.Println(file)

	v, _ := file.Stat()
	fmt.Println(v.Name(), v.Size())

	fmt.Println("------LEER ARCHIVO------")
	data := make([]byte, 800)

	c, err := file.Read(data)
	if err != nil {
		fmt.Println(err)
		os.Exit(1) //0 exitoso, 1 fallido
	}
	fmt.Println("bytes")
	fmt.Println(data[:c], c)
	fmt.Println("string")
	fmt.Println(string(data[:c]))
	fmt.Printf("read: %d - Bytes: %q \n", c, data[:c])

	//obtener y setear variables de entorno
	fmt.Println(os.Getenv("USERNAME"))
	fmt.Println(os.Getenv("MI_ENV"))
	os.Setenv("MI_ENV", "my value")
	fmt.Println(os.Getenv("MI_ENV"))

	//obtener direccion donde se encuentra
	dir, _ := os.Getwd()
	fmt.Println(dir)

	//Variable de entorno vacia y otra que no existe se imprimen igual
	os.Setenv("ENV_EXISTS", "")
	fmt.Println(os.Getenv("ENV_EXISTS"))
	fmt.Println(os.Getenv("EN_NOT_EXISTS"))

	//Esta funcion obtiene un booleano con base en la existencia de la variable de entorno
	env, ok := os.LookupEnv("ENV_EXISTS")
	fmt.Println(env, ok)

	env2, ok := os.LookupEnv("ENV_NOT_EXISTS")
	fmt.Println(env2, ok)

	//para bases de datos
	fmt.Println("----BASES DE DATOS----")
	os.Setenv("DB_USERNAME", "natalia")
	os.Setenv("DB_PASSWORD", "mysuperpassword")
	os.Setenv("DB_HOST", "localhost1")
	os.Setenv("DB_PORT", "6548")
	os.Setenv("DB_NAME", "users")

	dbURL := os.ExpandEnv("mongodb://${DB_USERNAME}:${DB_PASSWORD}@$DB_HOST:$DB_PORT/$DB_NAME")
	fmt.Println(dbURL)
}
