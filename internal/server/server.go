package server

import (
	"github.com/gin-gonic/gin"
	"portfolio-arya-service/internal/delivery/http/handler"
	"portfolio-arya-service/internal/repository"
	"portfolio-arya-service/internal/usecase"
)

func Run() {
	r := gin.Default()

	// Init dependencies
	aboutRepo := repository.NewAboutRepository()
	aboutUC := usecase.NewAboutUsecase(aboutRepo)
	handler.NewAboutHandler(r, aboutUC)

	servicesRepo := repository.NewServicesRepository()
	servicesUC := usecase.NewServicesUsecase(servicesRepo)
	handler.NewServicesHandler(r, servicesUC)

	// Start server
	r.Run() // default port 8080
}
