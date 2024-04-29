package controllers

import "github.com/gin-gonic/gin"

func HandleHomeRequest(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "WELCOME HOME", "error": nil})
}
