package postgresql

import (
	"fmt"
	"time"

	"github.com/doyyan/kubernetes/cmd/app/config"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbconn *gorm.DB

// CreateDBConnection Connection gets connection of postgresql database
func CreateDBConnection(ctx context.Context, config config.Config, logger *logrus.Logger) error {
	host := config.DBHOST
	user := config.DBUSER
	pass := config.DBPASSWORD
	dbname := config.DBNAME
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%s", host, user, pass, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Error(err)
		return err
	}
	dbconn = db
	sqlDB, err := dbconn.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	return nil
}

// GetDB returns a pre-initialised DB connection
func GetDB() *gorm.DB {
	return dbconn
}
