package router

import (
	"github.com/anthonnygm/apiopportunitiesgo/handler"

	"github.com/gin-gonic/gin"
)

func configRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", handler.GetOpeningHandler)
		v1.GET("/openings", handler.GetAllOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
	}
}