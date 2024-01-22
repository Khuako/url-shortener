package shortener

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"os"

	"github.com/itchyny/base58-go"
)

func sha256Of(input string) []byte {
	algorithm := sha256.New()
	algorithm.Write([]byte(input))
	return algorithm.Sum(nil)
}
func base58Encoded(bytes []byte) string {
	encoding := base58.BitcoinEncoding
	encoded, e := encoding.Encode(bytes)
	if e != nil {
		fmt.Println(e.Error())
		os.Exit(1)
	}
	return string(encoded)
}

func GenerateShortLink(initialLink, userId string) string {
	urlHashBytes := sha256Of(initialLink + userId)
	generatedNumber := []byte(new(big.Int).SetBytes(urlHashBytes).String())
	finalString := base58Encoded(generatedNumber)
	return finalString[:8]

}
