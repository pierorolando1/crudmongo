package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/pierorolando1/crudmongo/database"
	"github.com/pierorolando1/crudmongo/helpers"
	"github.com/pierorolando1/crudmongo/models"
)

func Logincontroller(c *gin.Context) {
	// get the data from json body
	println("login")
	//get the json body
	var json models.Login
	// pase the json
	err := c.BindJSON(&json)

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})

		return
	}

	println(json.Username)
	println(json.Password)

	// check if the user exists
	user, err := database.FindByUsername(json.Username)
	if err != nil || user == nil {
		c.JSON(400, gin.H{
			"ok":    false,
			"error": "user not found or something was wrong",
		})
		return
	}

	// check if the password is correct
	if !helpers.CheckPasswordHash(json.Password, user.Password) {
		c.JSON(400, gin.H{
			"ok":    false,
			"error": "password is incorrect",
		})
		return
	}

	println(string(user.Password))

	token, exp, err := helpers.CreateToken(json.Username)

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(200, gin.H{
		"ok":          true,
		"token":       token,
		"tokenexpire": exp,
		"user": gin.H{
			"username": json.Username,
		},
	})
}

func Registercontroller(c *gin.Context) {
	// get the data from json body
	println("register")
	//get the json body
	var json models.Register
	// pase the json
	err := c.BindJSON(&json)

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})

		return
	}

	// check if the user exists
	userfind, err := database.FindByUsername(json.Username)
	if err != nil || userfind != nil {
		c.JSON(400, gin.H{
			"ok":    false,
			"error": "user already exists",
		})
		return
	}

	hashedPassword, _ := helpers.HashPassword(json.Password)

	usercreated, err := database.CreateUser(gin.H{
		"username":    json.Username,
		"password":    hashedPassword,
		"displayName": json.DisplayName,
	})

	println(usercreated)

	token, exp, _ := helpers.CreateToken(json.Username)

	c.JSON(200, gin.H{
		"ok":          true,
		"token":       token,
		"tokenexpire": exp,
		"user": gin.H{
			"username":    json.Username,
			"displayName": json.DisplayName,
		},
	})
}
