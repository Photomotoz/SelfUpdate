package main

import (
	"github.com/fsnotify/fsnotify"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

func watcher() {
	watcher, err := fsnotify.NewWatcher()

	errorCheck(err)
	defer watcher.Close()

	// Start listening for events.
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Has(fsnotify.Create) || event.Has(fsnotify.Write) {
					log.Println("Valid change detected in dir, evaluating....")
					checkFiles(event)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				errorCheck(err)
			}
		}
	}()

	// Add a path.
	err = watcher.Add(RunningDirectory)
	if err != nil {
		log.Fatal(err)
	}

	// Run until told to stop
	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)
	<-quitChannel
	//time for cleanup before exit
	log.Println("Adios! 2")
}

func checkFiles(evt fsnotify.Event) {

	log.Println(evt.Name)

	_, _, binName := parsePath(evt.Name)

	if binName == BinaryName {

		log.Println("MATCH FOUND, CHECK FOR SIG")
		//Check Sig here

		// TODO Move this out
		//Load new binary
		log.Println("Starting new process : ", evt.Name)
		cmd := exec.Command(evt.Name)
		err := cmd.Run()
		errorCheck(err)
		//s := make(chan os.Signal, 1)

		//signal.Notify(s, os.Interrupt)
		//signal.Notify(s, syscall.SIGTERM)
		//os.Exit(0)
	}
	//
	////f, err := os.Open(runningDirectory)
	////errorCheck(err)
	//
	//info, err := os.Stat(runningDirectory + "/" + binaryName)
	//errorCheck(err)
	//log.Println(info)
}
