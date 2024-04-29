package router

import (
	"github.com/IBHMM/jwtauth/controllers"
	"github.com/IBHMM/jwtauth/middleware"
	"github.com/gin-gonic/gin"
)

func HomeRouter(route *gin.Engine) {
	route.GET("/", middleware.AuthMiddleware(), controllers.HandleHomeRequest)
}
