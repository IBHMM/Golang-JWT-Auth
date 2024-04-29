package controllers

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/IBHMM/jwtauth/config"
	"github.com/IBHMM/jwtauth/model"
	"github.com/IBHMM/jwtauth/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func HandleSignup(ctx *gin.Context) {
	var Body SingupRequest

	if err := ctx.BindJSON(&Body); err != nil {
		ctx.JSON(400, gin.H{"message": "Bad request body", "error": err})
		return
	}

	hashedPassword, err := utils.HashPassword(Body.Password)
	if err != nil {
		ctx.JSON(500, gin.H{"message": "Internal server error, please try again later"})
		return
	}

	user := model.User{
		ID:       utils.RandomID(),
		Name:     Body.Name,
		Email:    Body.Email,
		Password: hashedPassword,
		Status:   utils.StatusActive,
	}

	user.Role.Role = utils.RoleUser
	user.Role.ID = user.ID
	if result := config.DB.Create(&user); result.Error != nil {
		log.Println(result.Error)
		ctx.JSON(500, gin.H{"error": result.Error})
		return
	}

	ctx.JSON(200, gin.H{"message": "User succesfully signed up", "user": user, "error": nil})
}

func HandleSignin(ctx *gin.Context) {
	var Body SignInRequest

	if err := ctx.BindJSON(&Body); err != nil {
		ctx.JSON(400, gin.H{"error": err, "message": "Invalid email or password"})
		return
	}

	var user model.User

	if result := config.DB.First(&user, "email= ?", Body.Email); result.Error != nil {
		ctx.JSON(404, gin.H{"message": "Please, Sign up First"})
		return
	}

	if !utils.CheckPasswordHash(Body.Password, user.Password) {
		ctx.JSON(404, gin.H{"message": "Invalid Cridentials"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		ctx.JSON(400, gin.H{"error": "Failed to create token"})
		return
	}

	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
	ctx.SetCookie("User", user.Email, 3600*24*30, "", "", false, true)

	ctx.JSON(200, gin.H{"error": nil, "message": "Succesfull"})
}
