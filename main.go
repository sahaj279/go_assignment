package main

import (
	"log"

	"github.com/sahaj279/go_assignment/app"
)

func main() {
	if err := app.Init(); err != nil {
		log.Println("error in init", err)
	}
}
