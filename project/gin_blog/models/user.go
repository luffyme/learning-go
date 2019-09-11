package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	Model

	Name 	string `json:"name"`
	Desc 	string `json:"desc"`
	State   int    `json:"state"`
}

func GetUser(id int) (*User, error) {
	var user User

	err := db.Where("id = ? AND state = ? ", id, 2).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &user, nil
}