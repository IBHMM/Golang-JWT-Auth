package router

import (
	"github.com/IBHMM/jwtauth/controllers"
	"github.com/gin-gonic/gin"
)

func UserRouter(route *gin.Engine) {

	//SIGN IN
	route.POST("/api/user/signin", controllers.HandleSignin)

	//SING UP
	route.POST("/api/user/signup", controllers.HandleSignup)

	//LOG UOT
	route.PATCH("/api/user/logout")

	//Delete
	route.DELETE("/api/user/delete")

}
