package ft

import "fmt"

//Warning display a highlighted message on terminal
func Warning(message string) {
	fmt.Println(ClrY, "==========================================================")
	fmt.Println(ClrY, "=                                                        =")
	fmt.Println(ClrY, "=                                                        =")
	fmt.Println(Clr0, "=", message)
	fmt.Println(ClrY, "=                                                        =")
	fmt.Println(ClrY, "=                                                        =")
	fmt.Println(ClrY, "==========================================================")
	fmt.Println(ClrW)
}
