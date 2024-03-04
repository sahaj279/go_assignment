package main

import (
	"log"

	"github.com/sahaj279/go_assignment/service"
)

func main() {
	if err := service.Init(); err != nil {
		log.Println(err)
	}
}
