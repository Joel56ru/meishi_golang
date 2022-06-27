package handler

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "meishi_golang/docs"
	"meishi_golang/pkg/service"
)

//Handler Struct handler
type Handler struct {
	services *service.Service
}

//VeryCuteHandler Implement handler
func VeryCuteHandler(s *service.Service) *Handler {
	return &Handler{services: s}
}

//RoutingInitialization routing
func (h *Handler) RoutingInitialization() *gin.Engine {
	router := gin.New()
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1 := router.Group("/api/v1")
	{
		user := v1.Group("/user")
		{
			user.POST("/sign-up")
		}
	}
	return router
}
