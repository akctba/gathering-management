package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
	Name     string
}

type Gathering struct {
	gorm.Model
	Name        string
	Description string
	Date        string
	Location    string
	CreatorID   uint
	Creator     User `gorm:"foreignKey:CreatorID"`
	Invitees    []Invitee
}

type Invitee struct {
	gorm.Model
	Email       string
	Name        string
	GatheringID uint
	Gathering   Gathering `gorm:"foreignKey:GatheringID"`
	FoodPlateID *uint
	FoodPlate   *FoodPlate `gorm:"foreignKey:FoodPlateID"`
	BeverageID  *uint
	Beverage    *Beverage `gorm:"foreignKey:BeverageID"`
	RSVPStatus  string    `gorm:"default:'pending'"` // Can be 'pending', 'accepted', or 'declined'
}

type FoodPlate struct {
	gorm.Model
	Name        string
	Description string
	GatheringID uint
	Gathering   Gathering `gorm:"foreignKey:GatheringID"`
}

type Beverage struct {
	gorm.Model
	Name        string
	Description string
	GatheringID uint
	Gathering   Gathering `gorm:"foreignKey:GatheringID"`
}
