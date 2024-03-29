package http

import (
	"github.com/0z0sk0/simple-metrika-app/track"
	"github.com/gin-gonic/gin"
)

func RegisterEndpoints(router *gin.Engine, useCase track.UseCase) {
	handler := CreateHandler(useCase)

	trackGroup := router.Group("/track")
	{
		trackGroup.POST("", handler.Create)
	}
}
