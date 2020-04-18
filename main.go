package main

import (
	"math/rand"
	"os"
	"time"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"muskdaily.com/docs"

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

	docs.SwaggerInfo.Title = "Musk Daily API Documentation"
	docs.SwaggerInfo.Description = "Simple API descriptions for musk daily API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:" + configuration.Port
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":" + configuration.Port)
}
