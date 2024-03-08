package repository

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
)

func NewPostgres(dsn string) (Repository, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(fmt.Sprintf("failed to connect to database: %v", err))
	}
	if err := db.Migrator().AutoMigrate(&Card{}); err != nil {
		return nil, err
	}
	if err := db.Migrator().AutoMigrate(&Card{}); err != nil {
		return nil, err
	}
	return Postgres{
		db: db,
	}, nil
}

type Postgres struct {
	db *gorm.DB
}

func (p Postgres) Find() (res []Card, err error) {
	err = p.db.Find(&res).Error
	if err != nil {
		return nil, err
	}
	return
}

func (p Postgres) Delete(filter CardFilter) (res []Card, err error) {
	tx := p.db.Clauses(clause.Returning{})
	if filter.ID != nil {
		err = tx.Where("id = ?", *filter.ID).Delete(&res).Error
	} else if filter.Name != nil {
		err = tx.Where("name = ?", *filter.Name).Delete(&res).Error
	} else {
		err = tx.Where("1 = 1").Delete(&res).Error
	}
	if err != nil {
		return nil, err
	}
	return
}

func (p Postgres) Create(name string, answer string) error {
	return p.db.Create(&Card{
		Name:   name,
		Answer: answer,
	}).Error
}
