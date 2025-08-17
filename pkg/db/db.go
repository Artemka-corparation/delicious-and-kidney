package db

import (
	"delicious-and-kidney/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Db struct {
	*gorm.DB
}

func NewDb(conf *configs.Config) (*Db, error) {
	db, err := gorm.Open(postgres.Open(conf.Db.Dns), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &Db{db}, nil
}
