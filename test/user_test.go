package test

import (
	"github.com/jabutech/simple-blog/helper"
	"github.com/jabutech/simple-blog/util"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Function setup for connection to database test
func setupTestDb() *gorm.DB {
	// Load config
	config, err := util.LoadConfig(".") // "." as location file app.env in root folder
	helper.FatalError("error load config: ", err)

	// Open connection to db
	db, err := gorm.Open(mysql.Open(config.DBSource), &gorm.Config{})
	helper.FatalError("Connection to database failed!, err: ", err)

	return db
}

// func setupRouter(db *gorm.DB) *gin.Engine {
// 	// Repository
// 	userRepository := user.NewRepository(db)

// 	// Service
// 	userService := user.NewService(userRepository)

// 	// Handler
// 	userHandler := handler.NewUserHandler(userService)

// 	// Router
// 	router := router.NewRouter(userHandler)

// 	return router
// }
