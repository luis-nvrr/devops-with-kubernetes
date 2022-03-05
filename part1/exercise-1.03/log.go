package main

import (
	"fmt"
	"math/rand"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func main() {
	for {
		fmt.Printf("%s: %s\n", time.Now().UTC().String(), GetRandomHash())
		time.Sleep(5 * time.Second)
	}
}

func GetRandomHash() string {
	rand.Seed(time.Now().UTC().UnixNano())
	length := 36
	b := make([]rune, length)
	flag := false
	for i := range b {
		for _, j := range []int{9, 18, 27} {
			if i == j {
				b[i] = '-'
				flag = true
				break
			}
		}
		if flag {
			flag = false
			continue
		}
		rnd := rand.Intn(len(letters))
		b[i] = letters[rnd]

	}
	return string(b)
}
