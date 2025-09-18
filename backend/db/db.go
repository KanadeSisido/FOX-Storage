package db

import (
	"fmt"
	"fox-storage/model"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {

	dsn := createDsn(os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"),os.Getenv("DB_NAME"))
	//dsn := createDsn("process", "charles_sanders_peirce","fox_storage")

	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db.AutoMigrate(
		&model.User{},
		&model.Item{},
		&model.Permission{},
		&model.FileStorage{},
	)

	if err != nil {
		return nil, err
	}

	return db, nil
}


func createDsn(username string, password string, dbname string) string {

	return fmt.Sprintf("%s:%s@tcp(db:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, dbname)
}
