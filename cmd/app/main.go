package main

import (
	"WBTech_L0/internal/app"
	"log"
)

func main() {
	if err := app.StartServer(); err != nil {
		log.Fatal(err)
	}
}
