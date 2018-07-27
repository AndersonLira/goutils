package str

import (
	"testing"

	"github.com/andersonlira/goutils/str"
)

func TestNormalize(t *testing.T) {
	n := str.Normalize("Sebastião")
	if "Sebastiao" != n {
		t.Errorf("Result should be %s, but %s", "Sebastiao", n)
	}

	n = str.Normalize("ÁÉÍÓÚáéíóúÂÊÎÔÛÃẼĨÕŨ")
	if "AEIOUaeiouAEIOUAEIOU" != n {
		t.Errorf("Result should be %s, but %s", "AEIOUaeiouAEIOUAEIOU", n)
	}

	n = str.Normalize("ÇçŚśŕŔ")
	if "CcSsrR" != n {
		t.Errorf("Result should be %s, but %s", "CcSsrR", n)
	}

}
