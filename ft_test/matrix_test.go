package ft_test

import (
	"github.com/andersonlira/goutils/ft"
)

func ExampleMatrix() {
	m := ft.Matrix{
		Width:     30,
		Separator: '*',
	}
	m.Line()
	m.Println("Hello")
	m.Line()

	m.Line()
	m.Println("Hello Beautiful World. You are amazing!")
	m.Line()

	// Output:
	// ******************************
	// * Hello                      *
	// ******************************
	// ******************************
	// * Hello Beautiful World. You *
	// ******************************

}
