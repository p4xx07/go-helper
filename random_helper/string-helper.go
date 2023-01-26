package random_helper

import (
	"math/rand"
	"time"
)

type Complexity string

const (
	AZ                         Complexity = "abcdefghijklmnopqrstuvwxyz"
	AZCaps                     Complexity = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	AZAndCaps                  Complexity = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Numbers                    Complexity = "0123456789"
	AZAndCapsAndNumbers        Complexity = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	AZAndCapsAndNumbersSymbols Complexity = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*"
)

func Generate(length int, complexity Complexity) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	for i := range b {
		b[i] = complexity[rand.Intn(len(complexity))]
	}
	return string(b)
}
