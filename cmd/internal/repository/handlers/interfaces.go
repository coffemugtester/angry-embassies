package handlers

import "github.com/gin-gonic/gin"

type EmbassyHandlerImpl interface {
	GetDocuments(ctx *gin.Context)
}
