package utils

import (
	"math/rand"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func init() {
	rand.Seed(time.Now().UnixNano())
}

func arrayOfString(minLen, maxLen, cap int) []string {
	array := make([]string, cap)
	for i := 0; i < cap; i++ {
		array[i] = ""
	}
	return array
}

func ArrayOfInt(min, max, cap int) []int {
	array := make([]int, cap)
	for i := 0; i < cap; i++ {
		array[i] = rand.Intn(max-min) + min
	}
	return array
}

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func ArrayOfString(min, max, cap int) []string {
	array := make([]string, cap)
	for i := 0; i < cap; i++ {
		len := rand.Intn(max-min) + min
		array[i] = RandStringRunes(len)
	}
	return array
}
