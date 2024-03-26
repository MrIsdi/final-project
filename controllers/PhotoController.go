package controllers

import (
	"final-project/helpers"
	"final-project/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func (idb *InDB) StorePhoto(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	_, _, _ = idb, contentType, userData

	Photo := models.Photo{}
	UserID := uint(userData["id"].(float64))

	if contentType == "application/json" {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = UserID
	err := idb.DB.Create(&Photo).Error
	if err != nil {
		result := gin.H{
			"status":  false,
			"message": "Photo unsuccessfully created",
		}
		c.JSON(http.StatusBadRequest, result)
	}
	result := gin.H{
		"status":  true,
		"message": "Photo successfully created",
		"data":    Photo,
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) IndexPhoto(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	_, _, _ = idb, contentType, userData

	Photo := models.Photo{}

	if contentType == "application/json" {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	err := idb.DB.Find(&Photo).Error
	if err != nil {
		result := gin.H{
			"status":  false,
			"message": "Photo unsuccessfully view",
		}
		c.JSON(http.StatusBadRequest, result)
	}
	result := gin.H{
		"status":  true,
		"message": "Photo successfully view",
		"data":    Photo,
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) UpdatePhoto(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	_, _, _ = idb, contentType, userData

	Photo := models.Photo{}
	id, _ := strconv.Atoi(c.Param("id"))

	if contentType == "application/json" {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.ID = uint(id)
	err := idb.DB.Model(&Photo).Where("id = ?", id).Updates(models.Photo{Title: Photo.Title, Caption: Photo.Caption, PhotoUrl: Photo.PhotoUrl}).Error
	if err != nil {
		result := gin.H{
			"status":  false,
			"message": "Photo not found",
		}
		c.JSON(http.StatusNotFound, result)
		return
	}
	result := gin.H{
		"status":  true,
		"message": "Photo successfully updated",
		"data":    Photo,
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) DeletePhoto(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	_, _, _ = idb, contentType, userData

	Photo := models.Photo{}
	id := c.Param("id")

	if contentType == "application/json" {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	err := idb.DB.First(&Photo, id).Error
	if err != nil {
		result := gin.H{
			"status":  false,
			"message": "Photo not found",
		}
		c.JSON(http.StatusNotFound, result)
		return
	}
	err = idb.DB.Delete(&Photo).Error
	if err != nil {
		result := gin.H{
			"status":  false,
			"message": "Photo unsuccessfully deleted",
		}
		c.JSON(http.StatusBadRequest, result)
		return
	}
	result := gin.H{
		"status":  true,
		"message": "Photo successfully deleted",
	}
	c.JSON(http.StatusOK, result)
}
