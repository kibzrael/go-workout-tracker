package main

import (
	"kibzrael/workouttracker/cmd/data"
	"kibzrael/workouttracker/cmd/workouttracker"
)

func main() {
	data.InitDB()
	workouttracker.Router()
}