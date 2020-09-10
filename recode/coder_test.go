package gorecode

import (
	"fmt"
	"testing"
)

func TestInit(t *testing.T){
	_, err := MakeCoder("./_template.go")
	if err != nil  {
		t.Error("error should be nil here")
	}
}

func ExampleAddAfterLine(){
	coder, _ := MakeCoder("./_template.go")
	coder.AddAfterLine("import (",")","\t\"net/http\"")

	fmt.Println(coder.NewCodeContent())

	// Output:
	// package gorecode
	// 
	// import (
	// 	"fmt"
	// 	"errors"
	//	"net/http"
	// )
	//
	// func AnyFunction() {
	// 	fmt.Println("Any")
	// 	errors.New("New Error")
	// }	
}