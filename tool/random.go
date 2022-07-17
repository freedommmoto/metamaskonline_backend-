package tool

import (
	"context"
	"database/sql"
	db "github.com/freedommmoto/metamaskonline_backend/model/sqlc"
	"log"
	"math/rand"
	"strings"
)

const CodeNumber = "0123456789"

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomCodeNumber(long int, Queries *db.Queries) string {
	retry := 6
	i := 1
	newCode := ""
	for i < retry {
		i++

		newCode = RandomFormInput(CodeNumber, long)
		codeInBase, err := Queries.SelectCodeActive(context.Background(), newCode)
		if err != nil {
			if err == sql.ErrNoRows && len([]rune(newCode)) == long {
				i = 6
				return newCode
			}
			log.Println("error SelectCodeActive ", err)
		}
		if codeInBase != newCode && len([]rune(newCode)) == long {
			i = 6
			return newCode
		}
	}
	log.Println("error code is duplicate event retry ", retry)
	return newCode
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
