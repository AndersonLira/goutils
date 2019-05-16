package io_test

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"testing"
	"time"
	"github.com/andersonlira/goutils/io"
)

func TestWriteAndReadFile(t *testing.T) {
	fileName := getFileName()
	content := "Content Test"
	err := io.WriteFile(fileName, content)

	if err != nil {
		t.Error("Error was not expect here, but", err)
	}

	contentRead, err := io.ReadFile(fileName)

	if err != nil {
		t.Error("Error was not expect here, but", err)
	}

	if content != contentRead {
		t.Errorf("Content should be %s, but %s", content, contentRead)
	}
}

func TestWriteAndReadFilef(t *testing.T) {
	fileName := getFileName()
	content := "Content Test %s %d"
	err := io.WriteFilef(fileName, content, "A", 1)

	if err != nil {
		t.Error("Error was not expect here, but", err)
	}
	content = "Content Test A 1"
	contentRead, err := io.ReadFile(fileName)

	if err != nil {
		t.Error("Error was not expect here, but", err)
	}

	if content != contentRead {
		t.Errorf("Content should be %s, but %s", content, contentRead)
	}
}

func TestAppendFile(t *testing.T) {

	fileName := getFileName()

	content := "Content"
	err := io.AppendFile(fileName, content)

	contentRead, err := io.ReadFile(fileName)

	if err != nil {
		t.Error("Error was not expect here, but", err)
	}
	if !strings.Contains(contentRead, content) {
		t.Errorf("Content should be %s, but %s", content, contentRead)
	}

}

func TestAppendFilef(t *testing.T) {

	fileName := getFileName()
	content := "Content Test %s %d"
	io.WriteFile(fileName, "")

	io.AppendFilef(fileName, content, "A", 1)
	io.AppendFilef(fileName, content, "B", 2)

	contentRead, _ := io.ReadFile(fileName)

	if "Content Test A 1Content Test B 2" != contentRead {
		t.Errorf("Content should be 'Content Test A 1Content Test B 2', but %s", contentRead)
	}
}

func getFileName() (string){
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	prefix := os.TempDir()
	fileName := fmt.Sprintf("%s%s%dTestAppendFile.txt",prefix,string(os.PathSeparator), r.Int())
	return fileName
}