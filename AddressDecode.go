package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/ripemd160"
)

/*func P2PK() { //deprecated, but needed to decode early bitcoin blocks
	input, _ := hex.DecodeString("4104678afdb0fe5548271967f1a67130b7105cd6a828e03909a67962e0ea1f61deb649f6bc3f4cef38c4f35504e51ec112de5c384df7ba0b8d578a4c702b6bf11d5fac")
}

func OutputTypeIdentifier() {
	input_P2PK, _ := hex.DecodeString("")
	switch {
	case input[0:1] == 0x76:
		return //p2pkh
	case 81 <= input[0:1] <= 96:
		return //p2multisig
	case input[0:1] == 169:
		return //p2sh
	case input[0:1] == 106:
		return // null

	}
}*/

func HASH160() {
	input, _ := hex.DecodeString("4104678afdb0fe5548271967f1a67130b7105cd6a828e03909a67962e0ea1f61deb649f6bc3f4cef38c4f35504e51ec112de5c384df7ba0b8d578a4c702b6bf11d5fac")

	shaHasher := sha256.New()
	shaHasher.Write(input)
	sharesult := sha256.Sum256(input)

	ripeHasher := ripemd160.New()
	ripeHasher.Write(sharesult[:])

	result := ripeHasher.Sum(nil)
	fmt.Printf("%x", result)
	//   (x00 +  HASH160 (   x04 + (128bitPubKey)   )) + DOUBLE SHA CHECKSUM 4 bytes => base58
}
