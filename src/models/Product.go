package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key"`
	Name        string    `gorm:"type:varchar(100); not null"`
	Description string    `gorm:"type:varchar(255)"`
	Price       float64   `gorm:"type:float; not null"`
	Stock       int       `gorm:"type:int; not null"`
	Slug        string    `gorm:"type:varchar(100);uniqueIndex;not null"`
	CreatedAt   time.Time `gorm:"not null;default:now()"`
	UpdatedAt   time.Time `gorm:"not null;default:now()"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = uuid.New()
	p.Slug = generateSlug(p.Name)
	return
}
