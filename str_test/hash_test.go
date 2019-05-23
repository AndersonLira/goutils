package str

import (
	"testing"

	"github.com/andersonlira/goutils/str"
)

func TestHash(t *testing.T) {
	stringA := "Hello World"
	expectedA := uint32(3012568359)
	hashA := str.Hash(stringA)

	if hashA != expectedA {
		t.Errorf("Hash should be %d but %d",expectedA,hashA)
	}

	
	stringB := `
	Hello World
	Golang is fantastic... 
	`
	expectedB := uint32(1925370511)
	hashB := str.Hash(stringB)

	if hashB != expectedB {
		t.Errorf("Hash should be %d but %d",expectedB,hashB)
	}

}
