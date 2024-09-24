package main

import (
	"fmt"
	"log" //similar a fmt peroo imprime con un formato predeterminado, indicando fecha, hora y resultado a imprimir
	"os"
	"time"
)

func main() {
	log.Println("test")

	//log.Fatal("my error")
	//log.Panic("Error panic")
	date := time.Now()

	file, err := os.Create(fmt.Sprintf("%d.log", date.UnixNano()))
	if err != nil {
		log.Panic(err.Error())
	}
	//flag1 es fecha y hora y el segundo es el nombre del archivo con linea de impresion
	l := log.New(file, "success: ", log.LstdFlags|log.Lshortfile)

	l.Println("test1")
	l.Println("test2")
	l.Println("test3")

	l.SetPrefix("errors: ")
	l.Println("test1")
	l.Println("test2")
	l.Println("test3")
}
