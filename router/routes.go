package router

import (
	"github.com/gin-gonic/gin"
	"github.com/juliopragana/goestudo/handler/opening"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", opening.ShowOpeningHandler)

		v1.POST("/opening", opening.CreateOpeningHandler)

		v1.DELETE("/opening", opening.DeleteOpeningHandler)

		v1.PUT("/opening", opening.UpdateOpeningHandler)

		v1.GET("/openings", opening.ListOpeningsHandler)
	}
}
