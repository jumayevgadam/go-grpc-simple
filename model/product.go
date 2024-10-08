package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	Id          string    `gorm:"type:uuid;primaryKey" json:"id"`
	Name        string    `gorm:"type:varchar" json:"name"`
	Price       int64     `gorm:"type:bigint" json:"price"`
	IsAvailable bool      `gorm:"type:boolean" json:"is_available"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (base *Product) BeforeCreate(tx *gorm.DB) error {
	uuid := uuid.New()
	time := time.Now()
	tx.Statement.SetColumn("Id", uuid)
	tx.Statement.SetColumn("CreatedAt", time)
	tx.Statement.SetColumn("UpdatedAt", time)

	return nil
}

func (base *Product) BeforeUpdate(tx *gorm.DB) error {
	time := time.Now()
	tx.Statement.SetColumn("UpdatedAt", time)

	return nil
}
