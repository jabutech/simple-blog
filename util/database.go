package util

import (
	"github.com/jabutech/simple-blog/helper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Function setup for connection to database test
func SetupDb() *gorm.DB {
	// Load config
	config, err := LoadConfig(".") // "." as location file app.env in root folder
	helper.FatalError("error load config: ", err)

	// Open connection to db
	db, err := gorm.Open(mysql.Open(config.DBSource), &gorm.Config{})
	helper.FatalError("Connection to database failed!, err: ", err)

	return db
}

// Function setup for connection to database test
func SetupTestDb() *gorm.DB {
	// Load config
	config, err := LoadConfig("../") // "." as location file app.env in root folder
	helper.FatalError("error load config: ", err)

	// Open connection to db
	db, err := gorm.Open(mysql.Open(config.DBSourceTest), &gorm.Config{})
	helper.FatalError("Connection to database failed!, err: ", err)

	return db
}
