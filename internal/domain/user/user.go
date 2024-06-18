package user

import (
	"time"
)

type User struct {
	ID        string     `gorm:"type:uuid;primaryKey" json:"id"`
	Name      string     `gorm:"type:varchar(100)" json:"name"`
	Email     string     `gorm:"type:varchar(100);uniqueIndex" json:"email"`
	Password  string     `gorm:"type:varchar(100)" json:"password"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at"`
}
