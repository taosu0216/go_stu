package main

import (
	"JWT/router"
	"log"
)

func main() {
	r := router.Router()
	err := r.Run("0.0.0.0:23456")
	if err != nil {
		log.Fatalln("err is : ", err)
	}
}
