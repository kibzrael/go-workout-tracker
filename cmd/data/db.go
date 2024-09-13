package data

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)


func DB() *gorm.DB{
	db, err := gorm.Open(sqlite.Open("workouttracker.db"), &gorm.Config{TranslateError: true,})
	if err != nil {
		panic("failed to connect database")
	}
	return db;
}

func InitDB(){
	db := DB()

	db.AutoMigrate(&User{}, &Exercise{}, &Workout{}, &WorkoutActivity{})
}