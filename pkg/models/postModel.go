package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title   string
	Body    string
	User_id int
}
