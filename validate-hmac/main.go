package main

import (
	"crypto/sha256"
	"fmt"
	hmac "github.com/alexellis/hmac/v2"
)

func main() {
	input := []byte("input message from API")
	secret := []byte("so secret")

	digest := hmac.Sign(input, secret, sha256.New)
	fmt.Printf("Digest: %x\n", digest)

	err := hmac.Validate(input, fmt.Sprintf("sha256=%x", digest), string(secret))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Digest validated.\n")
}
