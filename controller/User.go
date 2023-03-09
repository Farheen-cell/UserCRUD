package controller

import (
	"fmt"
	"net/http"

	"github.com/Farheen-cell/UserCRUD/model"
	"github.com/gin-gonic/gin"
)

//GetUser .....get all user

func GetUser(c *gin.Context) {
	var user model.User
	err := model.GetAllUsers(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)

	} else {
		c.JSON(http.StatusOK, user)
	}
}

//CreateUser... Create User

func CreatUser(c *gin.Context) {
	var user model.User
	c.BindJSON(&user)
	err := model.CreateUser(&user)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//GetUserById...Get the user by Id

func GetUserById(c *gin.Context) {
	id := c.Params.ByName("id")
	var user model.User
	err := model.GetUserById(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// UpdateUser ....Update the user information
func UpdateUser(c *gin.Context) {
	var user model.User
	id := c.Params.ByName("id")
	err := model.GetUserById(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	c.BindJSON(&user)
	err2 := model.UpdateUser(&user, id)
	if err2 != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//DeleteUser....Delete the user

func DeleteUser(c *gin.Context) {
	var user model.User
	id := c.Params.ByName("id")
	err := model.DeleteUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
