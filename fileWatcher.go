package main

import (
	"github.com/fsnotify/fsnotify"
	"log"
)

func watcher() {
	log.Println("Starting watcher...")
	watcher, err := fsnotify.NewWatcher()

	errorCheck(err)
	defer func() {
		watcher.Close()
		replaceProcess()
	}()

	quit := make(chan bool)

	go func() {
		for {
			select {
			case <-quit:
				return
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Has(fsnotify.Create) || event.Has(fsnotify.Write) {
					log.Println("Valid change detected in dir, evaluating....")
					quit <- checkFiles(event)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				errorCheck(err)
			}
		}
	}()

	err = watcher.Add(RunningDirectory)
	errorCheck(err)
	<-quit

	log.Println("Shutting down watcher...")
}

func checkFiles(evt fsnotify.Event) bool {

	log.Println(evt.Name)
	_, _, binName := parsePath(evt.Name)

	if binName == BinaryName {
		log.Println("Valid file match, checking signature....")
		//TODO Check Sig here
		return true
	}
	log.Println("File or signature invalid.")
	return false
}
