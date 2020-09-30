package utils

import (
	"log"
)

func Err(err interface{}){
	if err != nil{
		log.Fatal(err)
	}
}


func generateRandomToken(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
