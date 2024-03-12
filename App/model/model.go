package model

import "gorm.io/gorm"

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&Role{}, &People{}, &Exercise{}, &Workout{})
}

func InitDefaultRoles(db *gorm.DB) {
	var role []Role
	value := db.Where("name=?", "ADMIN").First(&role)

	if value.RowsAffected == 0 {
		AdminRole := Role{Name: "ADMIN"}
		UserRole := Role{Name: "USER"}
		db.Create(&AdminRole)
		db.Create(&UserRole)
	}
	db.Find(&role)
}

type Role struct {
	ID      uint   `gorm:"primaryKey"`
	Name    string `gorm:"unique"`
	Peoples []People
}

type People struct {
	ID       uint `gorm:"primaryKey"`
	Name     string
	Age      int
	RoleID   uint
	Workouts []Workout
}

type Workout struct {
	ID        uint `gorm:"primaryKey"`
	Workout   string
	Exercises []Exercise
	PeopleID  uint
}

type Exercise struct {
	ID        uint `gorm:"primaryKey"`
	Exercise  string
	WorkoutID uint
}