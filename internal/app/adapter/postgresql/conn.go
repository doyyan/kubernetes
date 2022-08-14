package postgresql

import (
	"fmt"
	"time"

	"github.com/doyyan/kubernetes/cmd/app/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbconn *gorm.DB

// CreateDBConnection Connection gets connection of postgresql database
func CreateDBConnection(config config.Config) {
	host := config.DBHOST
	user := config.DBUSER
	pass := config.DBPASSWORD
	dbname := config.DBNAME
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%s", host, user, pass, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	dbconn = db

	sqlDB, err := dbconn.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
}

// GetDB returns a pre-initialised DB connection
func GetDB() *gorm.DB {
	return dbconn
}
