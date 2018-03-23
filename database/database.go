package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"monika/model"
)

type Database struct {
	db *gorm.DB
}

func NewDatabase(url string) (*Database, error) {
	db, err := gorm.Open("mysql", url)
	if err != nil {
		return nil, err
	}
	return &Database{db: db}, nil
}

func (d *Database) CreateTable(models ...interface{}) error {
	for _, v := range models {
		err := d.db.CreateTable(v).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *Database) InsertWeibo(weibo model.Weibo) error {
	err := d.db.Create(&weibo).Error
	return err
}

func (d *Database) InsertWeiboForward(weibo model.WeiboForward) error {
	err := d.db.Create(&weibo).Error
	return err
}

func (d *Database) Close() {
	d.db.Close()
}
