package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var getLogRequestMessage string

func main() {
	http.HandleFunc("/", handleLog)
	done := make(chan bool)
	go http.ListenAndServe(":3000", nil)

	go func() {
		for {
			fmt.Println(getMessage())
			time.Sleep(5 * time.Second)
		}
	}()
	<-done
}

func getMessage() string {
	return fmt.Sprintf("%s: %s\n", time.Now().UTC().String(), GetRandomHash())
}

func handleLog(w http.ResponseWriter, r *http.Request) {
	if getLogRequestMessage == "" {
		getLogRequestMessage = getMessage()
	}
	fmt.Fprint(w, getLogRequestMessage)
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
