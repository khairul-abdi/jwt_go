package controllers

import (
	"fmt"
	"myGram/database"
	"myGram/helpers"
	"myGram/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	db := database.GetDB()
	User := models.User{}
	c.ShouldBindJSON(&User)
	fmt.Println("USER => ", User)

	if User.Email == "" || User.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid Request",
			"message": "Username or Email required",
		})
		return
	}

	err := db.Debug().Create(&User).Error

	if err != nil {
		fmt.Println("ERROR")
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	fmt.Println(User)

	c.JSON(http.StatusCreated, gin.H{
		"id":    User.ID,
		"email": User.Email,
	})

}

func UserLogin(c *gin.Context) {
	db := database.GetDB()
	User := models.User{}
	password := ""

	c.ShouldBindJSON(&User)

	password = User.Password

	err := db.Debug().Where("email = ?", User.Email).Take(&User).Error

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email/password",
		})
		return
	}

	comparePass := helpers.ComparePass([]byte(User.Password), []byte(password))

	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email/password",
		})
		return
	}

	token := helpers.GenerateToken(User.ID, User.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
