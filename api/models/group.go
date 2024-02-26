package models

import "github.com/muhammetburakgolec/InvestHub-Backend/db"

// group_id,
// group_name,
// group_members

type Group struct {
	GroupId      uint   `json:"group_id"`
	GroupName    string `json:"group_name"`
	GroupMembers string `json:"group_members"`
}

func (g *Group) GetByGroupID(id uint) error {
	gormDB := db.GetDB()

	if err := gormDB.Where("group_id = ?", id).First(&g).Error; err != nil {
		return err
	}
	return nil
}

func (g *Group) CreateGroup() error {
	gormDB := db.GetDB()

	if err := gormDB.Where("group_id = ?", g.GroupId).First(&g).Error; err == nil {
		return nil
	}

	if err := gormDB.Create(g).Error; err != nil {
		return err
	}

	return nil
}
