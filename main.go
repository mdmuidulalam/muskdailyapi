package main

import (
	gin "github.com/gin-gonic/gin"
	config "muskdaily.com/config"
	data "muskdaily.com/data"
	manager "muskdaily.com/manager"
	routes "muskdaily.com/routes"
)

func main() {
	configuration := config.GetConfiguration()

	r := gin.Default()

	routes.Account{
		Base: routes.Base{R: r.Group("/account")},
		SignUpManager: manager.AccountManager{
			AccountData: &data.AccountData{},
		},
	}.New()

	r.Run(":" + configuration.Port)
}
