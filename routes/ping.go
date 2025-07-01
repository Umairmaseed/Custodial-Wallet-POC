package routes

import (
	"github.com/PrismFX-Dev/PrismFX_Wallet/controllers"
	"github.com/gin-gonic/gin"
)

func PingRoute(router *gin.RouterGroup) {
	auth := router.Group("/ping")
	{
		auth.GET(
			"",
			controllers.Ping,
		)
	}
}
