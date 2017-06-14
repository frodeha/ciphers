package main

import (
	"fmt"
)

type Cipher interface {
	Encrypt(s string) (string, error)
	Decrypt(s string) (string, error)
}

func main() {
	subCipher := NewSubstitutionCipher()
	cipherTest(subCipher)
}

func cipherTest(c Cipher) {
	s := "this is the text to encrypt"
	out, err := c.Encrypt(s)
	if err != nil {
		fmt.Printf("Error while encrypting: %v\n", err)
	}
	fmt.Printf("Encrypted text '%s'\n", out)

	out, err = c.Decrypt(out)
	if err != nil {
		fmt.Printf("Error while decrypting: %v\n", err)
	}

	fmt.Printf("Decrypted text '%s'\n", out)
}
