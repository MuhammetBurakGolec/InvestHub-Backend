package models

import (
	"github.com/muhammetburakgolec/InvestHub-Backend/db"
)

type User struct {
	ID         uint   `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	GroupID    uint   `json:"group_id" gorm:"column:group_id"` // GORM etiketi ile ilişkilendirilmiş
	IsAdmin    bool   `json:"is_admin"`
	IsInvestor bool   `json:"is_investor"`
	IsStudent  bool   `json:"is_student"`
}

func (u *User) FindByUsername(username string) error {
	gormDB := db.GetDB()

	if err := gormDB.Where("username = ?", username).First(&u).Error; err != nil {
		return err
	}
	return nil
}

func (u *User) GetByID(id uint) error {
	gormDB := db.GetDB()

	if err := gormDB.Where("id = ?", id).First(&u).Error; err != nil {
		return err
	}
	return nil
}
