package str

import "testing"

func TestNormalize(t *testing.T) {
	n := Normalize("Sebastião")
	if "Sebastiao" != n {
		t.Errorf("Result should be %s, but %s", "Sebastiao", n)
	}

	n = Normalize("ÁÉÍÓÚáéíóúÂÊÎÔÛÃẼĨÕŨ")
	if "AEIOUaeiouAEIOUAEIOU" != n {
		t.Errorf("Result should be %s, but %s", "AEIOUaeiouAEIOUAEIOU", n)
	}

	n = Normalize("ÇçŚśŕŔ")
	if "CcSsrR" != n {
		t.Errorf("Result should be %s, but %s", "CcSsrR", n)
	}

}
