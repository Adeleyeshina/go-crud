package controllers

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/AdeleyeShina/go-crud/initializer"
	"github.com/AdeleyeShina/go-crud/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllUser(ctx *gin.Context) {
	var users []models.User

	if err := initializer.DB.Find(&users).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"data":     "null",
			"messsage": err.Error(),
			"status":   "failed",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data":     users,
		"messsage": "User response",
		"status":   "success",
	})
}

func GetSingleUser(ctx *gin.Context) {
	var user models.User
	id := ctx.Param("id")

	pid, err := strconv.Atoi(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Id"})
		return
	}

	if err := initializer.DB.First(&user, "id =?", pid).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "User not Found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func CreateUser(ctx *gin.Context) {
	var body models.User
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if len(body.Password) < 6 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Password cannot be less than 6 characters"})
		return
	}

	user := models.User{Email: body.Email, Password: body.Password, Age: body.Age}
	if err := initializer.DB.Create(&user).Error; err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			ctx.JSON(http.StatusConflict, gin.H{"message": "Email Already Exist"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"data":     user,
		"messsage": "User Created",
		"status":   "success",
	})
}

func UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")
	var body models.User
	var user models.User

	pid, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Id type",
		})
		return
	}

	if err := initializer.DB.First(&user, pid).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "No User Found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// payload := models.User{Email: body.Email, Password: body.Password}
	user.Email = body.Email
	user.Password = body.Password

	if err := initializer.DB.Updates(&user).Error; err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			ctx.JSON(http.StatusConflict, gin.H{"message": "Email exist"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	var user models.User

	pid, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Id type",
		})
		return
	}

	if err := initializer.DB.First(&user, pid).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "No User Found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	if err := initializer.DB.Delete(&user).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Deleted"})

}
