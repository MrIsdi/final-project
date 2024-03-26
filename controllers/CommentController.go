package controllers

import (
	"final-project/helpers"
	"final-project/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func (idb *InDB) StoreComment(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	_, _, _ = idb, contentType, userData

	Comment := models.Comment{}
	UserID := uint(userData["id"].(float64))

	if contentType == "application/json" {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserID = UserID
	err := idb.DB.Create(&Comment).Error
	if err != nil {
		result := gin.H{
			"status":  false,
			"message": "Comment unsuccessfully created",
		}
		c.JSON(http.StatusBadRequest, result)
	}
	result := gin.H{
		"status":  true,
		"message": "Comment successfully created",
		"data":    Comment,
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) IndexComment(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	_, _, _ = idb, contentType, userData

	Comment := models.Comment{}

	if contentType == "application/json" {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	err := idb.DB.Find(&Comment).Error
	if err != nil {
		result := gin.H{
			"status":  false,
			"message": "Comment unsuccessfully view",
		}
		c.JSON(http.StatusBadRequest, result)
	}
	result := gin.H{
		"status":  true,
		"message": "Comment successfully view",
		"data":    Comment,
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) UpdateComment(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	_, _, _ = idb, contentType, userData

	Comment := models.Comment{}
	id, _ := strconv.Atoi(c.Param("id"))

	if contentType == "application/json" {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.ID = uint(id)
	err := idb.DB.Model(&Comment).Where("id = ?", id).Updates(models.Comment{Message: Comment.Message}).Error
	if err != nil {
		result := gin.H{
			"status":  false,
			"message": "Comment not found",
		}
		c.JSON(http.StatusNotFound, result)
		return
	}
	result := gin.H{
		"status":  true,
		"message": "Comment successfully updated",
		"data":    Comment,
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) DeleteComment(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	_, _, _ = idb, contentType, userData

	Comment := models.Comment{}
	id := c.Param("id")

	if contentType == "application/json" {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	err := idb.DB.First(&Comment, id).Error
	if err != nil {
		result := gin.H{
			"status":  false,
			"message": "Comment not found",
		}
		c.JSON(http.StatusNotFound, result)
		return
	}
	err = idb.DB.Delete(&Comment).Error
	if err != nil {
		result := gin.H{
			"status":  false,
			"message": "Comment unsuccessfully deleted",
		}
		c.JSON(http.StatusBadRequest, result)
		return
	}
	result := gin.H{
		"status":  true,
		"message": "Comment successfully deleted",
	}
	c.JSON(http.StatusOK, result)
}
