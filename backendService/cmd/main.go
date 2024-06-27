package main

import (
	"github.com/yakuzzaa/GoDone/backendService/internal/storage"
	"log"
)

func main() {
	storage.Connect()
	log.Println("Database connected successfully")
}
