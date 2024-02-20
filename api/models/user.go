package models

import (
	"fmt"
	"github.com/muhammetburakgolec/InvestHub-Backend/db"
)

type User struct {
	Id         uint   `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	GroupId    uint   `json:"group_id" gorm:"column:group_id"` // GORM etiketi ile ilişkilendirilmiş
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

func (u *User) Register() error {
	gormDB := db.GetDB()

	var count int64
	if err := gormDB.Model(&User{}).Where("username = ?", u.Username).Count(&count).Error; err != nil {

		return err
	}

	if count > 0 {

		return fmt.Errorf("username '%v' is already in use", u.Username)
	}

	if u.GroupId == 0 {
		u.GroupId = 1
	}

	if err := gormDB.Create(u).Error; err != nil {
		return err
	}

	return nil
}
