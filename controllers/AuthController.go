package controllers

import (
	"final-project/helpers"
	"final-project/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func (idb *InDB) Register(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	_, _ = idb, contentType

	User := models.User{}

	if contentType == "application/json" {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	err := idb.DB.Create(&User).Error
	if err != nil {
		result := gin.H{
			"status":  false,
			"message": "Account unsuccessfully created",
		}
		c.JSON(http.StatusOK, result)
	}
	result := gin.H{
		"status":  true,
		"message": "Account successfully created",
		"data":    User,
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) Login(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	_, _ = idb, contentType

	User := models.User{}
	password := ""

	if contentType == "application/json" {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	password = User.Password
	err := idb.DB.Debug().Where("email = ?", User.Email).Take(&User).Error

	if err != nil {
		result := gin.H{
			"status":  false,
			"message": "invalid email/password",
		}
		c.JSON(http.StatusUnauthorized, result)
		return
	}

	comparePass := helpers.ComparePass([]byte(User.Password), []byte(password))
	if !comparePass {
		result := gin.H{
			"status":  false,
			"message": "invalid email/password",
		}
		c.JSON(http.StatusUnauthorized, result)
		return
	}

	token := helpers.GenerateToken(User.ID, User.Email)
	result := gin.H{
		"status":  true,
		"message": "Account successfully logged in",
		"data":    token,
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) ChangeUser(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	_, _, _ = idb, contentType, userData

	User := models.User{}
	id, _ := strconv.Atoi(c.Param("id"))

	if contentType == "application/json" {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	User.ID = uint(id)
	err := idb.DB.Model(&User).Where("id = ?", id).Updates(models.User{Email: User.Email, Username: User.Username}).Error
	if err != nil {
		result := gin.H{
			"status":  false,
			"message": "Account not found",
		}
		c.JSON(http.StatusNotFound, result)
		return
	}
	result := gin.H{
		"status":  true,
		"message": "Account successfully updated",
		"data":    User,
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) DeleteUser(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	_, _, _ = idb, contentType, userData

	User := models.User{}
	id := c.Param("id")

	if contentType == "application/json" {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	err := idb.DB.First(&User, id).Error
	if err != nil {
		result := gin.H{
			"status":  false,
			"message": "Account not found",
		}
		c.JSON(http.StatusNotFound, result)
		return
	}
	err = idb.DB.Delete(&User).Error
	if err != nil {
		result := gin.H{
			"status":  false,
			"message": "Account unsuccessfully deleted",
		}
		c.JSON(http.StatusBadRequest, result)
		return
	}
	result := gin.H{
		"status":  true,
		"message": "Account successfully deleted",
	}
	c.JSON(http.StatusOK, result)
}
