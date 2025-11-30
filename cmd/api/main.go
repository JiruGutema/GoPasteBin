package main

import (
	"log"
	"gopastbin/internal/server"
	"github.com/joho/godotenv"
)

func main() {
	srv := server.NewServer("3000")
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
}
