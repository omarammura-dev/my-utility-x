package models

import "time"

type Link struct{
	Name string 
	Url string `binding:"required"`
	CreatedAt time.Time
	User_id int64
}