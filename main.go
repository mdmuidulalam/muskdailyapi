package main

import (
	"math/rand"
	"os"
	"time"

	gin "github.com/gin-gonic/gin"
	config "muskdaily.com/config"
	data "muskdaily.com/data"
	environmentEnum "muskdaily.com/enums/environment"
	manager "muskdaily.com/manager"
	routes "muskdaily.com/routes"
	service "muskdaily.com/service"
)

func main() {
	configuration := config.GetConfiguration()

	r := gin.Default()

	if configuration.Environment == environmentEnum.Live {
		myfile, _ := os.Create("panic.log")
		r.Use(gin.RecoveryWithWriter(myfile))
	}

	baseService := service.Service{
		Configuration: configuration,
	}

	baseData := data.Data{
		Configuration: configuration,
	}

	routes.Account{
		Base: routes.Base{R: r.Group("/account")},
		SignUpManager: manager.AccountManager{
			Manager:     manager.Manager{},
			AccountData: &data.AccountData{Data: baseData},
			EmailService: service.EmailService{
				Service: baseService,
			},
			RandomService: service.RandomService{
				Service:    baseService,
				SeededRand: rand.New(rand.NewSource(time.Now().UnixNano() + configuration.RandomSeedOffset)),
			},
		},
	}.New()

	r.Run(":" + configuration.Port)
}
