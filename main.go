package main

import (
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

	signedMessage := RSASignature.SignInput(input)

	signedJsonMessage, _ := json.Marshal(signedMessage)

	fmt.Println(string(signedJsonMessage))
}
