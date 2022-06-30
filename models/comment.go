package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	ID         int    `gorm:"not null" json:"id"`
	Contents   string `gorm:"not null" json:"contents"`
}

type CommentReport struct {
	gorm.Model
	ID         int    `gorm:"not null" json:"id"`
	Reason   string `gorm:"not null" json:"reason"`
	CommentId int `gorm:"not null" json:"comment_id"`
}