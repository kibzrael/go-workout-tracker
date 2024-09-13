package data

import (
	"time"

	"gorm.io/gorm"
)

type User struct{
	gorm.Model
	Email string `gorm:"unique; not null"`
	Password string
	LoggedAt *time.Time
}

type Workout struct{
	gorm.Model
	Title string
	User User `gorm:"foreignKey:ID"`
	activity []WorkoutActivity `gorm:"foreignKey:ID"`
}

type WorkoutActivity struct{
	gorm.Model
	weight uint `gorm:"default:0"`
	reps uint `gorm:"default:0"`
	progress uint `gorm:"default:0"`
	Exercise Exercise `gorm:"foreignKey:ID"`
}

type Exercise struct{
	gorm.Model
	Name string
	Description *string
	Category string
	Muscle string
}
