package handlers

import "github.com/gin-gonic/gin"

type Handlers struct {
}

func (h *Handlers) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up")
		auth.POST("/sign-in")

	}
	//api := router.Group("/api")

}
