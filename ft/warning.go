package ft

import "fmt"

//Warning display a highlighted message on terminal
func Warning(message ...interface{}) {
	fmt.Println(ClrY, "==========================================================")
	fmt.Println(ClrY, "=                                                        =")
	fmt.Println(ClrY, "=                                                        =")
	for _, it := range message {
		fmt.Println(Clr0, "=", it)
	}
	fmt.Println(ClrY, "=                                                        =")
	fmt.Println(ClrY, "=                                                        =")
	fmt.Println(ClrY, "==========================================================")
	fmt.Println(ClrW)
}
