package repo

import "gorm.io/gorm"

type TestDB struct {
	db *gorm.DB
}

type TestDBTransaction interface {
	RunWithRollback(func(tx *gorm.DB) error) error
}

func NewTestDB(db *gorm.DB) TestDBTransaction {
	return &TestDB{db: db}
}

func (t *TestDB) RunWithRollback(f func(tx *gorm.DB) error) error {
	tx := t.db.Begin()
	if err := f(tx); err != nil {
		return err
	}
	return tx.Rollback().Error
}
