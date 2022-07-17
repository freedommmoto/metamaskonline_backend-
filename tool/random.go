package tool

import (
	"math/rand"
	"strings"
)

const CodeNumber = "0123456789"

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomCodeNumber(long int) string {
	return RandomFormInput(CodeNumber, long)
}

func RandomFormInput(inputSting string, long int) string {
	var sb strings.Builder
	k := len(inputSting)
	for i := 0; i < long; i++ {
		c := inputSting[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}
