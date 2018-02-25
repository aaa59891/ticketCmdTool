package utils

import (
	"errors"

	"github.com/aaa59891/ticket/db"
	"github.com/jinzhu/gorm"
)

var (
	ErrNoTransaction = errors.New("This Transaction Is Empty.")
)

func Transactional(ts ...func(db2 *gorm.DB) error) error {
	if ts == nil || len(ts) == 0 {
		return ErrNoTransaction
	}
	tx := db.DB.Begin()

	for _, t := range ts {
		if err := t(tx); err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}
