package handlers

import "github.com/gin-gonic/gin"

type EmbassyHandlerImpl interface {
	GetDocument(ctx *gin.Context)
}
