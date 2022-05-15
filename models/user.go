package models

import "time"

type User struct {
	ID        uint      `json:"id"`
	FirstName string    `json:"firstName"`
	Surname   string    `json:"surname"`
	CreatedAt time.Time `json:"createdAt"`
}

type UserLoginModel struct {
	Username string `json:"username" xml:"username" form:"username"`
	Password string `json:"password" xml:"password" form:"password"`
}
