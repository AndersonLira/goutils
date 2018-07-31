package io_test

import (
	"testing"

	"github.com/andersonlira/goutils/io"
)

func TestWriteAndReadFile(t *testing.T) {
	fileName := "/tmp/201831071529TestWriteFile.txt"
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

func TestAppendFile(t *testing.T) {
	fileName := "/tmp/201831071529TestWriteFile.txt"
	content := ""
	io.WriteFile(fileName, content)

	content = "Content"
	err := io.AppendFile(fileName, content)

	contentRead, err := io.ReadFile(fileName)

	if content != contentRead {
		t.Errorf("Content should be %s, but %s", content, contentRead)
	}
	err = io.AppendFile(fileName, content)

	contentRead, err = io.ReadFile(fileName)

	if err != nil {
		t.Error("Error was not expect here, but", err)
	}
	if (content + content) != contentRead {
		t.Errorf("Content should be %s, but %s", content+content, contentRead)
	}

}
