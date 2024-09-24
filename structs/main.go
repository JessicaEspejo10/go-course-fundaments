package main

import (
	"encoding/json"
	"fmt"
	"fundaments/structs/commerce"
)

// la estructura puede contener los campos, sus tipos y opcionalmente los tags en `comillas invertidas`
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Address  string `json:"address,omitempty"`
	Age      uint8  `json:"age,omitempty"`
	LastName string `json:"last_name,omitempty"`
	//metodo en estructura
}

func (u User) Display() {
	v, _ := json.Marshal(u)
	fmt.Println(string(v))
}

// el * indica que se desea que la función afecte la estructura original
func (u *User) DefineName(name string) {
	u.Name = name
}

func main() {
	user := User{
		ID:   123,
		Name: "Nahuel",
	}

	user.Display()

	user.DefineName("Azul")
	user.Display()

	p1 := commerce.Product{
		Name:  "Heladera",
		Price: 20000,
		Type: commerce.Type{
			Code:        "A",
			Description: "Electrodomestico",
		},
		Tags:  []string{"heladera", "freezer", "refrigerador"},
		Count: 5,
	}

	p2 := commerce.Product{
		Name:  "Naranja",
		Price: 50,
		Type: commerce.Type{
			Code:        "B",
			Description: "Alimento",
		},
		Tags:  []string{"fruta", "alimentp"},
		Count: 10,
	}

	p3 := commerce.Product{
		Name:  "Cortina",
		Price: 4000,
		Type: commerce.Type{
			Code:        "C",
			Description: "Hogar",
		},
		Tags:  []string{"persiana", "ventanas", "cortina"},
		Count: 4,
	}

	car := commerce.NewCar(11213)
	car.AddProducts(p1, p2, p3)

	fmt.Println("PRODUCTS CAR")
	fmt.Println("Total products: ", len(car.Products))
	fmt.Printf("Total car %2.f\n", car.Total())
	fmt.Println()
}
