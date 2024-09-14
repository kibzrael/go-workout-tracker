package main

import (
	"kibzrael/workouttracker/cmd/data"
	"kibzrael/workouttracker/cmd/workouttracker"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil{
		log.Fatal("Failed to load .env", err)
	}

	data.InitDB()
	workouttracker.Router()
}