package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Moviment struct {
	ID         uuid.UUID      `gorm:"type:uuid;primary_key"`
	Type       TipoMovimiento `gorm:"type:varchar(10);not null"`
	Amount     int64          `gorm:"not null"`
	Date       time.Time      `gorm:"not null;default:CURRENT_DATE"`
	IDProducts uuid.UUID      `gorm:"type:uuid;not null"`
	IDUser     uuid.UUID      `gorm:"type:uuid;not null"`

	Producto Product `gorm:"foreignKey:IDProducts;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Usuario  User    `gorm:"foreignKey:IDUser;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type TipoMovimiento string

const (
	Entrada TipoMovimiento = "Entrada"
	Salida  TipoMovimiento = "Salida"
)

func (m *Moviment) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New()
	m.Date = time.Now()
	return
}
