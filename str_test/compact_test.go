package str

import (
	"testing"

	"github.com/andersonlira/goutils/str"
)

func TestCompact(t *testing.T) {
	s := `
		test with spaces
			tabs and
			line breaks	
	`
	expected := `testwithspacestabsandlinebreaks`

	compacted := str.Compact(s)

	if expected != compacted {
		t.Errorf("Compacted should be %s but %s",expected,compacted)
	}
	
}
