package ref_test

import (
	"testing"

	"github.com/andersonlira/goutils/ref"
)

func TestGetFunctionName(t *testing.T) {

	thisName := ref.GetFunctionName(TestGetFunctionName)

	if "TestGetFunctionName" != thisName {
		t.Errorf("TestGetFunctionName expected, but %s", thisName)
	}

}
