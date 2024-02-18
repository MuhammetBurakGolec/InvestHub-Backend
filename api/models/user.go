package models

import (
	"github.com/muhammetburakgolec/InvestHub-Backend/db"
)

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *User) FindByUsername(username string) error {
	gormDB := db.GetDB()

	if err := gormDB.Where("username = ?", username).First(&u).Error; err != nil {
		return err
	}
	return nil
}
