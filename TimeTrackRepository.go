package main

import "github.com/jinzhu/gorm"

type TimeTrackRepository interface {
	PersistNew() error
}

type TimeTrackRepositoryImpl struct {
	pgDB *gorm.DB
}

var dbTT TimeTrackRepositoryImpl

func (r *TimeTrackRepositoryImpl) PersistNew(dbPg *gorm.DB, t TimeTrack) error {
	tx := dbPg.Begin()
	if err := tx.Create(t).Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
