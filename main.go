package main

import (
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"os"

	"github.com/userforbidden/s_senthivel.challenge/RSASignature"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("App requires one argument")
		os.Exit(0)
	}

	input := os.Args[1]

	if len(input) > 250{
		fmt.Println("Input length cannot be longer than 250 characters")
		os.Exit(0)
	}
	
	keyFile := RSASignature.GetKeyPath()
	fmt.Println(keyFile)

	privateKeybyte , err := RSASignature.GetKeyData(keyFile)

	var privateKey *rsa.PrivateKey

	if err != nil{
		fmt.Println("Creating New Key")
		privateKey, _ = RSASignature.CreateNewKey(keyFile); 
	} else {
		fmt.Println("Using Existing Key")
		privateKey, _ = RSASignature.DecodeKeyData(privateKeybyte)
	}

	

	signedMessage := RSASignature.SignInput(input,privateKey)

	signedJsonMessage, _ := json.Marshal(signedMessage)

	fmt.Println(string(signedJsonMessage))
}
