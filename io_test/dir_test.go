package io_test

import (
	"testing"
	"github.com/andersonlira/goutils/io"
)
func TestListFiles(t *testing.T) {
	files, err := io.ListFiles("../_support/", []string{})

	if err != nil {
		t.Errorf("Error was not expected here but %v",err)
	}


	if len(files) != 6 {
		t.Errorf("Files variable %v", files)
		t.Errorf("Size expected was 5, but %v",len(files))
	}

	filesTxt, _ := io.ListFiles("../_support/", []string{"txt"})

	if len(filesTxt) != 2 {
		t.Errorf("Files variable %v", filesTxt)
		t.Errorf("Size expected was 2, but %v",len(filesTxt))
	}

	filesTxtAndDat, _ := io.ListFiles("../_support/", []string{"txt","dat"})

	if len(filesTxtAndDat) != 4 {
		t.Errorf("Files variable %v", filesTxtAndDat)
		t.Errorf("Size expected was 4, but %v",len(filesTxtAndDat))
	}
}
