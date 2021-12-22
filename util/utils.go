package util

import (
	"math/rand"
	"time"
)

func RandomString(n int) string {
	var str = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	res := make([]byte, n)

	rand.Seed(time.Now().UnixNano())
	for i := range res {
		res[i] = str[rand.Intn(len(str))]
	}
	return string(res)
}
