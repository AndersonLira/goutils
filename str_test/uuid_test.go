package str

import (
	"testing"

	"github.com/andersonlira/goutils/str"
)

func TestNewUUID(t *testing.T) {
	size := 36
	n := str.NewUUID()

	if len(n) != size {
		t.Errorf("Size was wrong. It should be %d but %d", size, len(n))
	}
	//Test one million or more
	//times := 1000000
	times := 10000
	checkDuplicate := make(map[string]bool)

	for i := 0; i < times; i++ {
		key := str.NewUUID()
		checkDuplicate[key] = true
	}

	if len(checkDuplicate) != times {
		t.Errorf("Size should be %d but %d", times, len(checkDuplicate))
	}
}
