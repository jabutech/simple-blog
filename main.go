package main

import (
	"github.com/jabutech/simple-blog/handler"
	"github.com/jabutech/simple-blog/helper"
	"github.com/jabutech/simple-blog/router"
	"github.com/jabutech/simple-blog/util"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// Load config
	config, err := util.LoadConfig(".") // "." as location file app.env in root folder
	helper.FatalError("error load config: ", err)

	// Open connection to db
	_, err = gorm.Open(mysql.Open(config.DBSource), &gorm.Config{})
	helper.FatalError("Connection to database failed!, err: ", err)

	// Handler
	pingHandler := handler.NewPingHandler()

	// Router
	router := router.NewPingRouter(pingHandler)

	// Run router
	router.Run()
}
