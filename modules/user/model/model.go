package model

import (
	"time"

	"github.com/lucsky/cuid"
	"gorm.io/gorm"
)

type User struct {
	Id        string    `gorm:"primaryKey" json:"id"`
	Name      *string   `gorm:"type:varchar(300)" json:"name"`
	Email     *string   `gorm:"type:varchar(300)" json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.Id = cuid.New()
	return nil
}
