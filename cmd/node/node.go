package main

import (
	"log"

	"github.com/freitzzz/monica/internal/mq"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	ctx, err := mq.NewContext()

	if err != nil {
		log.Fatal(err)
	}

	s, err := mq.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	defer mq.Close(s)
	mq.RegisterPub(s)
}
