package main

import (
	"github.com/exPriceD/WBTech-L0_Microservice/internal/app"
	"log"
)

func main() {
	if err := app.StartServer(); err != nil {
		log.Fatal(err)
	}
}
