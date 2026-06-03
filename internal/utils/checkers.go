package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"

	"github.com/google/uuid"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func IsNotNull(s ...string) bool {
	for _, v := range s {
		if strings.TrimSpace(v) == "" {
			return false
		}
	}
	return true
}

func UUIDGenerator(s string) string {
	return fmt.Sprintf("%s-%s", s, uuid.New().String())
}

func GenerateOrder() string {
	n, _ := rand.Int(rand.Reader, big.NewInt(1000))

	l := ""
	for i := 0; i < 3; i++ {
		idx, _ := rand.Int(rand.Reader, big.NewInt(int64(len(letterBytes))))
		l += string(letterBytes[idx.Int64()])
	}

	return fmt.Sprintf("%s-%d", l, n.Int64()) // abc-012
}
