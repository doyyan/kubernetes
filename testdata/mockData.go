package testdata

import (
	"context"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/doyyan/kubernetes/internal/app/adapter/controller"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetDB() *gorm.DB {
	db1, _, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 db1,
		PreferSimpleProtocol: true,
	})
	db2, _ := gorm.Open(dialector, &gorm.Config{})
	return db2
}

func CreateController() controller.Controller {
	return controller.Controller{
		Context: context.Background(),
		Logger:  logrus.New(),
	}
}
