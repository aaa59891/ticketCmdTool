package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

type Ticket struct {
	Id        uint   `gorm:"primary_key"`
	Code      string `sql:"index"`
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func (ticket *Ticket) Create(tx *gorm.DB) error {
	return tx.Create(ticket).Error
}

func (ticket Ticket) String() string {
	return fmt.Sprintf("Code: %s\tEmail: %s\tUpdated: %s", ticket.Code, ticket.Email, ticket.UpdatedAt)
}

func (ticket Ticket) StringWithoutEmail() string {
	return fmt.Sprintf("Code: %s\tUpdated: %s", ticket.Code, ticket.UpdatedAt)
}
