package main

import (
	"os"
)

func main() {
	defer func() {
		if recover() != nil {
			os.Exit(1)
		}
	}()

	if len(os.Args) != 2 {
		os.Stderr.Write([]byte("Usage : ./hmac <input> <key>\n"))
		return
	}

	key := []byte(os.Args[2])
	hmac := Digest([]byte(os.Args[1]), key)

	os.Stdout.Write(hmac)
	os.Exit(0)
}