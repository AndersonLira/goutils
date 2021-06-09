package ft_test

import (
	"testing"
	"github.com/andersonlira/goutils/ft"
)

// func ExampleMatrix() {
// 	m := ft.Matrix{
// 		Width:     30,
// 		Separator: '*',
// 	}
// 	m.Line()
// 	m.Println("Hello")
// 	m.Line()

// 	m.Line()
// 	m.Println("Hello Beautiful World. You are amazing!")
// 	m.Line()

// 	// Output:
// 	// ******************************
// 	// * Hello                      *
// 	// ******************************
// 	// ******************************
// 	// * Hello Beautiful World. You *
// 	// ******************************

// }

func TestGetText(t *testing.T){
	m := ft.Matrix{
		Width: 5,
		Separator: '+',
		LineColor: ft.RED,
		TextColor: ft.BLUE,
	}

	expected := `+++++
+ AB+
+ 12+
+++++
`


	m.Line()
	m.Println("ABC")
	m.Println("12345678910")
	m.Line()

	if expected != m.GetText() {
		t.Error(m.GetText())
		//t.Errorf("Expected was %s but %s",expected,m.GetText())
	}
}