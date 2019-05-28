package watcher_test

import (
	"os"
	"testing"

	"github.com/andersonlira/goutils/watcher"
	"github.com/andersonlira/goutils/io"
)

func TestStart(t *testing.T){
	paths := []string {"./"}
	wf := watcher.WatcherFile{Paths:paths}


	wf.Start()
	go io.WriteFile("./test.txt","content")

	fileChanged := <-wf.FileChanged
	if fileChanged != "test.txt" {
		t.Errorf("File changed should be test.txt but %s",fileChanged)
	}
	fileChanged = <-wf.FileChanged
	if fileChanged != "test.txt" {
		t.Errorf("File changed should be test.txt but %s",fileChanged)
	}

	go os.Remove("./test.txt")

	fileChanged = <-wf.FileChanged
	if fileChanged != "test.txt" {
		t.Errorf("File changed should be test.txt but %s",fileChanged)
	}

}