package db

import "gorm.io/gorm"

type TransactionMaker struct {
	tx *gorm.DB
}

func (t *TransactionMaker) Close(err *error) {
	if *err != nil {
		t.tx.Rollback()
		return
	}
	t.tx.Commit()
}
