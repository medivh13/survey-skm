package models

/*
 * Author      : Jody (jody.almaida@gmail)
 * Modifier    :
 * Domain      : auth-skm
 */

import (
	"time"
)

type Users struct {
	ID        int64     `gorm:"id"`
	Name      string    `gorm:"display_name"`
	Email     string    `gorm:"email_address"`
	Password  string    `gorm:"password"`
	CreatedAt time.Time `gorm:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at"`
}
