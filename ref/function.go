package ref

import (
	"reflect"
	"runtime"
	"strings"
)

//GetFunctionName returns name from giving string
func GetFunctionName(i interface{}) string {
	longName := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	names := strings.Split(longName, ".")
	if len(names) > 0 {
		return names[len(names)-1]
	}
	return ""
}
