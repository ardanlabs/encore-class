package main

import "log"

func main() {
	if err := GenKey(); err != nil {
		log.Fatal(err)
	}
}

// GenKey creates an x509 private/public key for auth tokens.
func GenKey() error {
	return nil
}
