package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id        int            `gorm:"primaryKey;autoIncrement" json:"id"`
	Title     string         `gorm:"type:varchar(50);not null"  json:"title"`
	Email     string         `gorm:"type:varchar(100);uniqueIndex;not null" json:"email"`
	FirstName string         `gorm:"type:varchar(100);not null" json:"firstName"`
	LastName  string         `gorm:"type:varchar(100);not null" json:"lastName"`
	Gender    int            `gorm:"not null" json:"gender"`
	Dob       time.Time      `gorm:"type:date;not null" json:"dob"`
	IsDeleted bool           `gorm:"default:false" json:"is_deleted"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
