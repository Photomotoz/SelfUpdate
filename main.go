package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

// TODO - Pull in through ENV VARS
const publicKey = ""
const privateKey = ""
const version string = "0.0.1"

var RunningDirectory = ""
var BinaryName = ""

func main() {
	log.Println("Starting SelfUpdate version " + version)

	// Store public key
	//blockPub, _ := pem.Decode([]byte(publicKey))
	//pubKey, _ := x509.ParsePKCS1PublicKey(blockPub.Bytes)

	//blockPriv, _ := pem.Decode([]byte(privateKey))
	//privKey, _ := x509.ParsePKCS1PrivateKey(blockPriv.Bytes)

	exe, err := os.Executable()
	errorCheck(err)
	_, RunningDirectory, BinaryName = parsePath(exe)

	log.Println("Running path : " + RunningDirectory)
	log.Println("Running binary : " + BinaryName)

	go watcher()

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
	<-exit

	log.Println("Shutting down SelfUpdate version " + version)
}

func replaceProcess() {
	log.Println("Starting new process...")
	err := syscall.Exec(RunningDirectory+"/"+BinaryName, nil, os.Environ())
	if err != nil {
		log.Println(err)
		go watcher()
	}
}
