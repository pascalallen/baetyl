package Crypto

import (
	"encoding/hex"
	"math/rand"
)

type Crypto string

func Generate() Crypto {
	bytes := make([]byte, 32)
	rand.Read(bytes)

	crypto := make([]byte, hex.EncodedLen(len(bytes)))
	hex.Encode(crypto, bytes)

	return Crypto(crypto)
}
