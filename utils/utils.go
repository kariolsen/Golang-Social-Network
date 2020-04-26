package utils

import (
	"log"
)

func Err(err interface{}){
	if err != nil{
		log.Fatal(err)
	}
}

