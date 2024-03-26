package controllers

import (
	"final-project/helpers"
	"final-project/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func (idb *InDB) StoreSocialMedia(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	_, _, _ = idb, contentType, userData

	SocialMedia := models.SocialMedia{}
	UserID := uint(userData["id"].(float64))

	if contentType == "application/json" {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	SocialMedia.UserID = UserID
	err := idb.DB.Create(&SocialMedia).Error
	if err != nil {
		result := gin.H{
			"status":  false,
			"message": "SocialMedia unsuccessfully created",
		}
		c.JSON(http.StatusBadRequest, result)
	}
	result := gin.H{
		"status":  true,
		"message": "SocialMedia successfully created",
		"data":    SocialMedia,
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) IndexSocialMedia(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	_, _, _ = idb, contentType, userData

	SocialMedia := models.SocialMedia{}

	if contentType == "application/json" {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	err := idb.DB.Find(&SocialMedia).Error
	if err != nil {
		result := gin.H{
			"status":  false,
			"message": "SocialMedia unsuccessfully view",
		}
		c.JSON(http.StatusBadRequest, result)
	}
	result := gin.H{
		"status":  true,
		"message": "SocialMedia successfully view",
		"data":    SocialMedia,
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) UpdateSocialMedia(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	_, _, _ = idb, contentType, userData

	SocialMedia := models.SocialMedia{}
	id, _ := strconv.Atoi(c.Param("id"))

	if contentType == "application/json" {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	SocialMedia.ID = uint(id)
	err := idb.DB.Model(&SocialMedia).Where("id = ?", id).Updates(models.SocialMedia{Name: SocialMedia.Name, SocialMediaUrl: SocialMedia.SocialMediaUrl}).Error
	if err != nil {
		result := gin.H{
			"status":  false,
			"message": "SocialMedia not found",
		}
		c.JSON(http.StatusNotFound, result)
		return
	}
	result := gin.H{
		"status":  true,
		"message": "SocialMedia successfully updated",
		"data":    SocialMedia,
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) DeleteSocialMedia(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	_, _, _ = idb, contentType, userData

	SocialMedia := models.SocialMedia{}
	id := c.Param("id")

	if contentType == "application/json" {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	err := idb.DB.First(&SocialMedia, id).Error
	if err != nil {
		result := gin.H{
			"status":  false,
			"message": "SocialMedia not found",
		}
		c.JSON(http.StatusNotFound, result)
		return
	}
	err = idb.DB.Delete(&SocialMedia).Error
	if err != nil {
		result := gin.H{
			"status":  false,
			"message": "SocialMedia unsuccessfully deleted",
		}
		c.JSON(http.StatusBadRequest, result)
		return
	}
	result := gin.H{
		"status":  true,
		"message": "SocialMedia successfully deleted",
	}
	c.JSON(http.StatusOK, result)
}
