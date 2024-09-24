package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	now := time.Now()
	p(now)

	then := time.Date(2024, 10, 8, 20, 34, 58, 654987, time.UTC)
	p(then)

	//agregar una hora
	then = then.Add(1 * time.Hour)
	p(then)

	//agregar un dia
	then = then.Add(24 * time.Hour)
	p(then)
	//agregar un media hora
	then = then.Add(30 * time.Minute)
	p(then)

	p("--------------------------------------")

	p("Then es antes que Now: ", then.Before(now))
	p("Then es despues que Now: ", then.After(now))
	p("Then es igual que Now: ", then.Equal(now))
}
