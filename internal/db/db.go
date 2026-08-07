package db

import (
	"gorm.io/gorm"

	"gorm.io/driver/sqlite"
)

type Db struct {
	Db *gorm.DB
}

func (db *Db) Open() (err error) {
	db.Db, err = gorm.Open(sqlite.Open("monci.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}

func (db *Db) Migrate() error {
	return db.Db.AutoMigrate(&Pipeline{}, &Job{}, &Step{})
}
