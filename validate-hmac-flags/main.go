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
	flag.StringVar(&modeVar, "mode", "", "define if the cli [generate] or [validate] the digest")
	flag.StringVar(&digestVar, "digest", "", "digest to validate")
	flag.Parse()

	if len(strings.TrimSpace(secretVar)) == 0 {
		panic("--secret is required")
	}
	if len(strings.TrimSpace(modeVar)) == 0 {
		panic("--mode is required")
	}

	digest := hmac.Sign([]byte(inputVar), []byte(secretVar), sha256.New)

	if modeVar == "generate" {
		fmt.Printf("Digest: %x\n", digest)
		return
	} else if modeVar == "validate" {
		err := hmac.Validate([]byte(inputVar), fmt.Sprintf("sha256=%x", digest), string(secretVar))
		if err != nil {
			panic(err)
		}
		fmt.Printf("Digest valid.")
	}
}
