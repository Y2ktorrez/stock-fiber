package models

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key"`
	Name      string    `gorm:"type:varchar(100); not null"`
	Email     string    `gorm:"type:varchar(100);uniqueIndex;not null"`
	Telefono  string    `gorm:"type:varchar(100)"`
	Rol       Role      `gorm:"type:varchar(100);default:Cliente"`
	Slug      string    `gorm:"type:varchar(100);uniqueIndex;not null"`
	CreatedAt time.Time `gorm:"not null;default:now()"`
	UpdatedAt time.Time `gorm:"not null;default:now()"`
}

type Role string

const (
	Cliente   Role = "Cliente"
	Dueño     Role = "Dueño"
	Proveedor Role = "Proveedor"
)

func generateSlug(name string) string {
	slug := strings.ToLower(strings.ReplaceAll(name, " ", "_"))
	return slug
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	u.Slug = generateSlug(u.Name)
	return
}
