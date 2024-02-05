package model

import (
	"github.com/lucsky/cuid"
	"gorm.io/gorm"
)

type User struct {
	Id    string `gorm:"primaryKey" json:"id"`
	Name  string `gorm:"type:varchar(300)" json:"name"`
	Email string `gorm:"type:varchar(300)" json:"email"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.Id = cuid.New()
	return nil
}
