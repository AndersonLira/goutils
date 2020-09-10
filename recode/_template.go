package gorecode

import (
	"fmt"
	"errors"
)

func AnyFunction() {
	fmt.Println("Any")
	errors.New("New Error")
}