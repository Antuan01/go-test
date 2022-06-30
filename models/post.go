package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	ID         int    `gorm:"not null" json:"id"`
	Contents   string `gorm:"not null" json:"contents"`
}

type PostReport struct {
	gorm.Model
	ID         int    `gorm:"not null" json:"id"`
	Reason   string `gorm:"not null" json:"reason"`
	PostId int `gorm:"not null" json:"post_id"`
}