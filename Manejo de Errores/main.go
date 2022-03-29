// You can edit this code!
// Click here and start typing.
package main

import (
	"errors"
	"fmt"
)

func division(a float32, b float32) (float32, error) {
	var err error = nil
	if b == 0 {
		err := errors.New("Division por cero")
		return 0, err
	}
	result := a / b
	return result, err
}

func main() {
	fmt.Println("Hello, 世界")

	div, err := division(6, 0)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(div)
}
