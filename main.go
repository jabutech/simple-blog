package main

import (
	"fmt"

	"github.com/jabutech/simple-blog/helper"
	"github.com/jabutech/simple-blog/util"
)

func main() {
	// Load config
	config, err := util.LoadConfig(".") // "." as location file app.env in root folder
	helper.FatalError(err)

	fmt.Println(config.DBSource)
}
