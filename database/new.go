package database

import (
	"fmt"
	"time"

	"github.com/capdale/was/config"
	"github.com/capdale/was/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func New(mysqlConfig *config.Mysql) (db *Database, err error) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/was?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlConfig.Username,
		mysqlConfig.Password,
		mysqlConfig.Address,
		mysqlConfig.Port,
	)

	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}

	sqldb, err := d.DB()
	if err != nil {
		return
	}

	sqldb.SetMaxIdleConns(10)
	sqldb.SetMaxOpenConns(0)
	sqldb.SetConnMaxLifetime(time.Minute * 3)

	db = &Database{
		DB: d,
	}
	err = db.AutoMigrate()

	return
}

func (d *Database) AutoMigrate() (err error) {
	err = d.DB.AutoMigrate(
		&model.User{},
		&model.Collection{},
		&model.ImageQueue{},
	)
	return err
}
