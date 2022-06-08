package util

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Function setup for connection to database test
func SetupDb() *gorm.DB {
	// Load config
	config, err := LoadConfig(".") // "." as location file app.env in root folder
	if err != nil {
		log.Fatal("error load config: ", err.Error())
	}

	// Open connection to db
	db, err := gorm.Open(mysql.Open(config.DBSource), &gorm.Config{})
	if err != nil {
		log.Fatal("Connection to database failed!, err: ", err.Error())
	}

	return db
}

// Function setup for connection to database test
func SetupTestDb() *gorm.DB {
	// Load config
	config, err := LoadConfig("../") // "." as location file app.env in root folder
	if err != nil {
		log.Fatal("error load config: ", err.Error())
	}

	// Open connection to db
	db, err := gorm.Open(mysql.Open(config.DBSourceTest), &gorm.Config{})
	if err != nil {
		log.Fatal("Connection to database failed!, err: ", err.Error())
	}

	return db
}
