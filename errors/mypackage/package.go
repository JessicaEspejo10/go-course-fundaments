package mypackage

import "fmt"

func Run() {

	defer func() {
		fmt.Println("end my run function")
	}()
	
	panic("panic in run function")
}
