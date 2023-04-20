package app

import (
	"fmt"
	"os"

	"github.com/ainmtsn1999/digitalent_final_hacktiv8/config"
	"github.com/ainmtsn1999/digitalent_final_hacktiv8/handlers"
	"github.com/ainmtsn1999/digitalent_final_hacktiv8/repositories"
	"github.com/ainmtsn1999/digitalent_final_hacktiv8/routes"
	"github.com/ainmtsn1999/digitalent_final_hacktiv8/services"
	"github.com/gin-gonic/gin"
)

var router = gin.New()

func StartApplication() {
	repo := repositories.NewRepo(config.GORM.DB)
	service := services.NewService(repo)
	server := handlers.NewHttpServer(service)

	routes.InitApi(router, server)

	port := os.Getenv("APP_PORT")
	router.Run(fmt.Sprintf(":%s", port))
}
