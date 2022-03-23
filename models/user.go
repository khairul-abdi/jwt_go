package models

import (
	"time"

	"myGram/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"not null;type:varchar(20);column:username;uniqueIndex" json:"username" validate:"required-Your username is required,username-Invalid username format"`
	Email     string    `gorm:"not null;type:varchar(100);column:email;uniqueIndex" json:"email" validate:"required-Your email is required,email-Invalid email format"`
	Password  string    `gorm:"type:varchar(100);column:password"`
	Age       int       `gorm:"column:age"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(u)
	if err != nil {
		return
	}

	u.Password = helpers.HashPass(u.Password)

	return
}
