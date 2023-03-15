package watcher_test

import (
	"os"
	"testing"

	"github.com/andersonlira/goutils/io"
	"github.com/andersonlira/goutils/watcher"
)

func TestStart(t *testing.T) {
	paths := []string{"./"}
	wf := watcher.WatcherFile{Paths: paths}

	wf.Start()
	io.WriteFile("./test.txt", "content")
	go io.WriteFile("./test.txt", "content")

	fileChanged := <-wf.FileChanged
	if fileChanged != "./test.txt" {
		t.Errorf("File changed should be ./test.txt but %s", fileChanged)
	}
	fileChanged = <-wf.FileChanged
	if fileChanged != "./test.txt" {
		t.Errorf("File changed should be ./test.txt but %s", fileChanged)
	}

	os.Remove("./test.txt")

}
