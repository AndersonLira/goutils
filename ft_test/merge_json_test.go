package ft_test

import (
	"testing"
	"github.com/andersonlira/goutils/ft"
	"github.com/andersonlira/goutils/io"

)


func TestMergeFiles(t *testing.T) {

	_,err := ft.MergeJson("","")
	if err == nil {
		t.Error("Error was expected here, but it was nil")
	}

	jsonWeak, _ := io.ReadFile("../_support_json/weak.json")
	jsonStrong, _ := io.ReadFile("../_support_json/strong.json")

	result, _ := ft.MergeJson(jsonWeak,jsonStrong)

	if len(result.NewKeys) != 4 {
		t.Errorf("NewKeys should be 4, but %v",len(result.NewKeys))
	}

	if len(result.MissingKeys) != 1 {
		t.Errorf("MissingKeys should be 1, but %v",len(result.MissingKeys))
	}
	

}