package ft

import (
	"fmt"
	"time"

	"github.com/andersonlira/goutils/str"
)

//Log prints
func Log(message interface{}) {
	fmt.Printf("%v - %v\n",str.FmtDate("dd/MM/yyyy HH:mm:ss",time.Now()), message)
}

