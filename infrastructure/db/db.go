package db

import (
	"fmt"
	"github.com/gabrielmq/codepix/domain/model"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	_ "gorm.io/driver/sqlite"
)

func init() {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	err := godotenv.Load(fmt.Sprintf("%s/../../.env", basepath))
	if err != nil {
		log.Fatalf("Error loading .env files")
	}
}

func ConnectDB(env string) *gorm.DB {
	var db *gorm.DB
	var err error

	if env != "test" {
		db, err = gorm.Open(os.Getenv("dbType"), os.Getenv("dsn"))
	} else {
		db, err = gorm.Open(os.Getenv("dbTypeTest"), os.Getenv("dsnTest"))
	}

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		panic(err)
	}

	if os.Getenv("debug") == "true" {
		db.LogMode(true)
	}

	if os.Getenv("AutoMigrateDb") == "true" {
		db.AutoMigrate(&model.Bank{}, &model.Account{}, &model.PixKey{}, &model.Transaction{})
	}

	return db
}
