package main

import (
	"errors"
	"fmt"
	"fundaments/errors/mypackage"
)

type MyCustomError struct {
	ID string
}

func (m MyCustomError) Error() string {
	return fmt.Sprintf("error with id: %s", m.ID)
}

func main() {
	var err error
	err = errors.New("My new error")
	fmt.Println(err)
	//imprimir como string
	fmt.Println(err.Error())

	err2 := fmt.Errorf("My format error, string: %s, number %d", "hola", 231)
	fmt.Println(err2)
	//imprimir como string
	fmt.Println(err2.Error())

	err3 := TestError(1)
	fmt.Println(err3)
	fmt.Println(errors.As(err3, &MyCustomError{}))

	err4 := TestError(4)
	fmt.Println(err4)
	fmt.Println(errors.As(err4, &MyCustomError{}))
	fmt.Println("-----------------------------------")
	err5 := errors.Join(err, err2, err3)
	fmt.Println(err5)

	fmt.Println(errors.Is(err5, err2))
	fmt.Println(errors.Is(err5, err4))

	//DEFEAR para manejo de errores, el defer debe estar arriba de cualquier error que pueda pasar
	defer func() {
		fmt.Println("end main")
		r := recover()
		if r != nil {
			fmt.Println("Recover in ", r)
		}
	}()

	/*
		v := 0
		_ = 5 / v
	*/
	//PANIC para forzar errores
	//panic("My panic message")
	mypackage.Run()

}

func TestError(v int) error {
	if v == 1 {
		return MyCustomError{"1254"}
	} else {
		return errors.New("My else error")
	}

}
