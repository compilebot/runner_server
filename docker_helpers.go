package main

import (
	"fmt"
	"math/rand"
)

func generateID(lang string) (id string) {
	ts := randomString(20 - len(lang))
	id = fmt.Sprintf("%s-runner-%v", lang, ts)
	id = id[:20]
	return
}

func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(97, 122))
	}
	return string(bytes)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
