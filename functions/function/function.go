package function

import (
	"errors"
	"fmt"
)

type Operation int

const (
	SUM Operation = iota
	SUB
	DIV
	MUL
	TEST
)

func Display(myValue int64) {
	fmt.Printf("type: %T, Value: %d, addres: %v \n", myValue, myValue, &myValue)
	fmt.Println()
}

// si no se especifica, toma el tipo de dato del ultimo argumento de la función
func Add(x, y int) int {
	return x + y
}

func RepeatString(increment int, value string) {
	for i := 0; i < increment; i++ {
		fmt.Print(value)

	}
	fmt.Println()
}

func Calc(op Operation, x, y float64) (float64, error) {
	switch op {
	case SUM:
		return x + y, nil
	case SUB:
		return x - y, nil
	case DIV:
		if y == 0 {
			return 0, errors.New("y mustn't be zero")
		} else {
			return x / y, nil
		}
	case MUL:
		return x * y, nil
	}
	return 0, errors.New("Operation not valid")
}

// se especifica el retorno en el inicio de la función, por eso no necesita especificar cual es el return
func Split(v int) (x, y int) {
	x = v * 4 / 9
	y = v - x
	return
}

func Msum(values ...float64) float64 {
	var sum float64
	for _, v := range values {
		sum += v
	}
	return sum
}

func MOperations(op Operation, values ...float64) (float64, error) {
	if len(values) == 0 {
		return 0, errors.New("There aren't values")
	}
	sum := values[0]
	for _, v := range values[1:] {
		switch op {
		case SUM:
			sum += v
		case SUB:
			sum -= v
		case DIV:
			if v == 0 {
				return 0, errors.New("Mustn't be zero")
			}
			sum /= v
		case MUL:
			sum *= v
		}
	}
	return sum, nil
}
