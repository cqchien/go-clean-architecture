package models

type User struct {
	ID       string `gorm:"primary_key"`
	Username string
	Password string
	Limit    int
}


