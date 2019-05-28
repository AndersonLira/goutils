package watcher

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
)

//WatcherFile is prepared to notify when some file changes in giving paths
type WatcherFile struct {
	//Paths that the watcher should watch
	Paths []string 
	FileChanged chan string
	watcher *fsnotify.Watcher
}



//WatcherFile is a observer 
func (wf *WatcherFile) Start(){
	wf.FileChanged = make(chan string)
	// creates a new file watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("ERROR", err)
	}
	wf.watcher = watcher

	//
	//done := make(chan bool)
	//
	go func() {
		for {
			select {
			// watch for events
			case event := <-watcher.Events:
				eventDetected(wf,event)
				// watch for errors
			case  <-watcher.Errors:
				wf.FileChanged <- "Error"
			}
		}
	}()

	// out of the box fsnotify can watch a single file, or a single directory
	for _,item := range wf.Paths {
		if err := watcher.Add(item); err != nil {
			fmt.Println("ERROR", err)
		}
	}

	//<-done
}

func eventDetected(wf *WatcherFile,event fsnotify.Event){
	wf.FileChanged <- event.Name
}