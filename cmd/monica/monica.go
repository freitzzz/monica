package main

import (
	"github.com/freitzzz/monica/internal/logging"
	_ "github.com/joho/godotenv/autoload"
	as "github.com/palavrapasse/aspirador/pkg"
)

func main() {
	logging.Aspirador = as.WithClients(logging.CreateAspiradorClients(""))

	logging.Aspirador.Info("hello world")
}
