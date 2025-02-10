package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
	"strings"

	hmac "github.com/alexellis/hmac/v2"
)

func main() {
	var inputVar string
	var secretVar string
	var digestVar string
	var modeVar string

	flag.StringVar(&inputVar, "message", "", "message to create a digest from")
	flag.StringVar(&secretVar, "secret", "", "secret for the digest")
	flag.StringVar(&modeVar, "mode", "", "define if the cli generate or validate the digest")
	flag.StringVar(&digestVar, "digest", "", "digest to validate")
	flag.Parse()

	if len(strings.TrimSpace(secretVar)) == 0 {
		panic("--secret is required")
	}
	if len(strings.TrimSpace(secretVar)) == 0 && (modeVar != "generate" || modeVar != "validate") {
		panic("--secret is required")
	}

	fmt.Printf("Computing hash for: %q\n", inputVar)
	fmt.Printf("Secret:%q\n", secretVar)

	digest := hmac.Sign([]byte(inputVar), []byte(secretVar), sha256.New)

	fmt.Printf("Digest: %x\n", digest)
}
