package main

import (
	"github.com/freitzzz/monica/internal/logging"
	"github.com/freitzzz/monica/internal/mq"
	_ "github.com/joho/godotenv/autoload"
	as "github.com/palavrapasse/aspirador/pkg"
	"github.com/pebbe/zmq4"
)

func main() {
	logging.Aspirador = as.WithClients(logging.CreateAspiradorClients(mq.ServerAddress()))

	ctx, err := zmq4.NewContext()

	if err != nil {
		panic(err)
	}

	s, err := mq.Start(ctx)

	if err != nil {
		panic(err)
	}

	mq.RegisterHandlers(s)
}
