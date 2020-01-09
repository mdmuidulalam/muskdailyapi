package main

import (
	"os"

	gin "github.com/gin-gonic/gin"
	config "muskdaily.com/config"
	data "muskdaily.com/data"
	environmentEnum "muskdaily.com/enums"
	manager "muskdaily.com/manager"
	routes "muskdaily.com/routes"
)

func main() {
	configuration := config.GetConfiguration()

	r := gin.Default()

	if configuration.Environment == environmentEnum.Live {
		myfile, _ := os.Create("panic.log")
		r.Use(gin.RecoveryWithWriter(myfile))
	}

	routes.Account{
		Base: routes.Base{R: r.Group("/account")},
		SignUpManager: manager.AccountManager{
			AccountData: &data.AccountData{},
		},
	}.New()

	r.Run(":" + configuration.Port)
}
