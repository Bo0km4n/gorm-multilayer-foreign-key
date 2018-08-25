package main

import "time"

type Region struct {
	ID int
	// Shops     []Shop `gorm:"foreginkey:RegionID"`
	Shops     []Shop `gorm:"foreignkey:ID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Shop struct {
	ID int
	// RegionID  int
	Name      string
	Books     []Book `gorm:"foreignkey:ShopID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Book struct {
	ID        int
	ShopID    int
	Name      string
	Price     int
	CreatedAt time.Time
	UpdatedAt time.Time
}
