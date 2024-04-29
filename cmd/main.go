package main

import (
	"fmt"

	"github.com/IBHMM/jwtauth/config"
	"github.com/IBHMM/jwtauth/router"
	"github.com/gin-gonic/gin"
)

func init() {
	variables := &config.Variables{}
	if err := variables.Load(); err != nil {
		fmt.Println(err)
	}
	err := config.Connect()
	if err == nil {
		fmt.Println(err)
	}
}

func main() {
	engine := gin.Default()

	router.UserRouter(engine)
	router.HomeRouter(engine)

	engine.Run()
}
