package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

// TODO - Pull in through ENV VARS
const publicKey string = "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAtQYVoep7bctIeYn6PUrs\n6D8e4xCrLHut60xqgSXYja9tpBCHrM5JPjVKWfCORNGrB6lRREzfg51L8QzGxdK9\nKVeSttpg+cikemt0VCr2+Nhq7JN+j3H1PRjwXSLlmup0IzaStWxZSlE2IMmROeSK\nsy72W1aVlYN4i4dhgXQKH2dvMmjB9Tu6+PdivwVnMk0KfutdA52ZzlCxCUObeZjx\nVxx0Eqn+Dkj+mRZRUVuEffUH/8REXs7pBYautfJC5g0ocx/4x11ZwpsMkqh0daN/\nEEL+F37GHK9tJv/BQGhTpUXfrIwusSZs1hWxokA2f98OqsL9sBRnpbBkksshtGcj\n2wIDAQAB\n-----END PUBLIC KEY-----"
const privateKey = "-----BEGIN PRIVATE KEY-----\nMIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQC1BhWh6ntty0h5\nifo9SuzoPx7jEKsse63rTGqBJdiNr22kEIeszkk+NUpZ8I5E0asHqVFETN+DnUvx\nDMbF0r0pV5K22mD5yKR6a3RUKvb42Grsk36PcfU9GPBdIuWa6nQjNpK1bFlKUTYg\nyZE55IqzLvZbVpWVg3iLh2GBdAofZ28yaMH1O7r492K/BWcyTQp+610DnZnOULEJ\nQ5t5mPFXHHQSqf4OSP6ZFlFRW4R99Qf/xERezukFhq618kLmDShzH/jHXVnCmwyS\nqHR1o38QQv4XfsYcr20m/8FAaFOlRd+sjC6xJmzWFbGiQDZ/3w6qwv2wFGelsGSS\nyyG0ZyPbAgMBAAECggEAEB8RGSDlSiNi+nP3xbKhA5Pvf8BRxj3jgMOG7qZ8sS2Q\n4Z4ZPyIF8tTTIMxkFyvnA+En3CHOfA2roDWHlEGO2Bo09jTfkw8z2rUsIHzNDt9T\nMmt2R3caJI1OhrTfZtr/ac1n5mFWWGCAJg5dcR7x3UBO71Z4Lt8NArwVSm6XojEx\nJSxZLOVvkOsIWDvK19Eq0eeceIV0tMtk7mZmbh6rucNrB6tNkOUKS6RTW+KrzlZ7\nUU6RivRdnzB1tqq9TwUFq8QzWdsgUt2Z5k+tHEk+s8E8hDouNQ63C/9VBfKykCHD\nJvxghcKvtSlP7nUzCRZKqZOd7aUGV1xK772gMZAfqQKBgQDZMyTTcLq+nX/ZnALe\nmfIUr4MTvr8UHqE3uMi3k14RsTpcmEvpOFaOsiCbabLQuTKuhC0JREYpJZE4amAl\nZAlRMZ8y4MrAYzQmDD7LBvoZt5oY8GlTVHeeDDKf5No183gm01z6+9drvJF3Tlzb\nCbSZZour8kSiBBMevkKwS9LdaQKBgQDVXI9WQx9wRM7Uzw1dNf6aWt6LGZSrtBxw\nWogk26IFjCx9la3KQwbj7k0mND2VQyJ3UXHHhsoUFlgRYqTpQcli8YEpo+GyOIWM\nXulpvxa+fuNZ5Ilg1QdyWXD/WjUp36s78NMdEchDDX9a572og67ShcygAGVitLA6\n4pwjzvuaowKBgQDAC5ulJQjJoAM6itpyrpXAuotgewkCpi6QUV+QCcpkRFtBIhwE\nw56tH0z7fZLxQL9SVWZoyihrG19zDYOjq/cdBo4n82DyiQTEG5Gt/KZI9PFY2cGX\nP5lgcpKUlEpo57S2R67v8JvCFAJfBjfOKJ+5/1TCHT2YL5gkV8Apl3GD+QKBgQCB\nv43zGVKPBWj6BQfNfT5dQ5E96cp9OERrFsLgFyhKU1ni7RkIfQY9qfcEmpQwpujY\nwpB8k71jxcq8l85NgxIit1aiqeRgavrbyAQNqIT+R0+epNneJu4mXbnXr1XRJxGM\nnADmFqBAFn8yGuc9CQdRQsiTLGp7QgJBbPki1YYwHQKBgQDVQfn0EugMUR/KgZMT\n7e4A8XDWtJ7uzGWCx6lyLXNtjXLSdwXytFyiHXioKGWp9cZSIbyzzcprjjgWiIqE\nj3dUgTWUUilcU5AILcLnUV4VSN8cOVpfapdwNff7PaI3bFFJJ6Q167cPIa/Z0pUz\neVMHYbTwfoXvUoPxxd6b2cyzjw==\n-----END PRIVATE KEY-----"

const version string = "0.0.1"

var RunningDirectory string = ""
var BinaryName string = ""

// Application watches the folder from which it is running for changes
// If a new binary is detected then we check it's signature and see if we can decode

// Watch for binary to change, if it has, check it's signature,
// If that matches then accept the new binary, and load it,
// Transfer all running process to it, and kill this process
func main() {
	log.Println("Starting SelfUpdate version " + version)

	//// Store public key
	//blockPub, _ := pem.Decode([]byte(publicKey))
	//pubKey, _ := x509.ParsePKCS1PublicKey(blockPub.Bytes)
	//
	//blockPriv, _ := pem.Decode([]byte(privateKey))
	//privKey, _ := x509.ParsePKCS1PrivateKey(blockPriv.Bytes)

	exe, err := os.Executable()
	errorCheck(err)
	_, RunningDirectory, BinaryName = parsePath(exe)

	log.Println("Running path : " + RunningDirectory)
	log.Println("Running binary : " + BinaryName)

	go watcher()

	// Run until told to stop
	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)
	<-quitChannel
	//time for cleanup before exit
	log.Println("Adios!")
}

// Wrap with decorator or something to check for errors? Bubble them up to a logger?
func verifySig() {

}
