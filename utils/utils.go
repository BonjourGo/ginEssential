package utils

import (
	"math/rand"
	"time"
)

func RandString(n int) string {
	var randString = []byte("gdsjfhskfhksdfsjkfhhjhjksdakjhkjHbndbklbnhpGSHDSJKDgisuGFFHVVCVJ")
	result := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = randString[rand.Intn(len(randString))]
	}
	return string(result)
}
