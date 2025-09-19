package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `gorm:"unique;not null" form:"username" json:"username" binding:"required,min=3,max=20"`
	Password string `gorm:"not null" form:"password" json:"password" binding:"min=6,max=20"`
	Email    string `gorm:"unique;not null" form:"email" json:"email" binding:"omitempty,email"`
}

type Post struct {
	gorm.Model
	Title   string `gorm:"not null" form:"title" json:"title" binding:"omitempty,min=3,max=100"`
	Content string `gorm:"type:text;not null" form:"content" json:"content" binding:"omitempty,min=1"`
	UserID  uint   `gorm:"not null"`
	User    User   `gorm:"foreignKey:UserID" binding:"omitempty"`
}

type Comment struct {
	gorm.Model
	Content string `gorm:"type:text;not null" form:"content" json:"content" binding:"omitempty,min=1"`
	PostID  uint   `gorm:"not null"  form:"postId" json:"postId"`
	Post    Post   `gorm:"foreignKey:PostID"  binding:"omitempty"`
	UserID  uint   `gorm:"not null"`
	User    User   `gorm:"foreignKey:UserID"  binding:"omitempty"`
}
