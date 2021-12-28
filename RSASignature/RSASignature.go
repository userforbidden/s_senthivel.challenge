package RSASignature

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

const keyFilename = "privatekey.pem"
/*
Declaring a struct with name RSAKey

which contains PrivateKey and Public key compliant to rsa format
*/

type RSAKey struct{
	PrivateKey *rsa.PrivateKey
	PublicKey *rsa.PublicKey
}

/* 
Declaring a Struct with name SignedJsonStructure 

which contains the message signature and Pubkey as per the Schema requirement
 */
 type SignedMessageStructure struct {
	 Message string `json:"message"`
	 Signature string `json:"signature"`
	 Pubkey string `json:"pubkey"`
 }

 /*
 	Given a input string and rsa.ProvateKey the input will be signed with RSA private and the 
	 SignedMessageStructure will be returned
 */

func SignInput(input string,privateKey *rsa.PrivateKey) SignedMessageStructure{
	
	hash := crypto.SHA256.New()
	hash.Write([]byte(input))
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey,crypto.SHA256,hash.Sum(nil))

	if err != nil{
		log.Fatalf("Unable to sign message : %v",err)
	}

	pemBytes := pem.EncodeToMemory(&pem.Block{
		Type: "PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(&privateKey.PublicKey),
	})

	return SignedMessageStructure{
		Message: input,
		Signature: base64.StdEncoding.EncodeToString(signature),
		Pubkey: string(pemBytes),
	}
}

/*
	This function get the directory where the privatekey.pem will be stored
	$HOME/.local/share/signer/privatekey.pem
	Used filepath package to make sure the directory names are portable between different OS  
*/
func GetKeyPath() string{
	home, err := os.UserHomeDir()

	// fmt.Println(home)

	if err != nil{
		log.Fatalf("Cannot determine the current Users home directory")
	}

	homeDir,_ := filepath.Abs(home)
	keyFilePath := filepath.Join(homeDir,".local","share","signer",keyFilename)

	return keyFilePath

}

/* 
	This function will get the contents of the privatekeyfile if exists
*/
func GetKeyData(filepath string) ([]byte,error){
	
	keyData, err := ioutil.ReadFile(filepath)
	if err != nil{
		return nil, err
	}
	return keyData, nil
	
}

func DecodeKeyData(keyData []byte) (*rsa.PrivateKey,error){
	
	// fmt.Println("Using existing Key file")
	block, _ := pem.Decode([]byte(keyData))
	if block == nil{
		return nil, fmt.Errorf("RSA private key format is incorrect")
	}
	
	privKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil{
		log.Fatalf("Key is not parsed: %v", err)
	}

	return privKey,nil
}

/*
	This function will create new key into the given filename
*/

func CreateNewKey(filename string)(*rsa.PrivateKey,error){
	err := os.MkdirAll(filepath.Dir(filename), os.ModePerm)

	if err != nil{
		log.Fatalf("could not make parent directories for file %s: %v", filename, err)
	}

	keyfile, err := os.Create(filename)
	if err != nil{
		log.Fatalf("file cannot be created %s: %v", filename, err)
	}
	defer keyfile.Close()
	
	privKey, err := rsa.GenerateKey(rand.Reader, 2056)
	if err != nil{
		log.Fatalf("Could not generate RSA Key: %v", err)
	}
	err = pem.Encode(keyfile,&pem.Block{
		Type: "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privKey),
	})
	if err != nil{
		log.Fatalf("Could not write key to file %s: %v", filename, err)
	}

	return privKey, nil

}
