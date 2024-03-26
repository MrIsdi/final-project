package routers

import (
	"final-project/controllers"
	"final-project/database"
	"final-project/middlewares"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func StartApp() *gin.Engine {
	db := database.DBInit()
	inDB := &controllers.InDB{DB: db}
	router := gin.Default()

	userAuth := router.Group("/auth")
	{
		userAuth.POST("/register", inDB.Register)
		userAuth.POST("/login", inDB.Login)
	}
	userChange := router.Group("/user")
	{
		userChange.Use(middlewares.Authentication())
		userChange.PUT("/:id", inDB.ChangeUser)
		userChange.DELETE("/:id", inDB.DeleteUser)
	}
	photoRouter := router.Group("/photo")
	{
		photoRouter.Use(middlewares.Authentication())
		photoRouter.POST("/", inDB.StorePhoto)
		photoRouter.GET("/", inDB.IndexPhoto)
		photoRouter.PUT("/:id", inDB.UpdatePhoto)
		photoRouter.DELETE("/:id", inDB.DeletePhoto)
	}
	commentRouter := router.Group("/comment")
	{
		commentRouter.Use(middlewares.Authentication())
		commentRouter.POST("/", inDB.StoreComment)
		commentRouter.GET("/", inDB.IndexComment)
		commentRouter.PUT("/:id", inDB.UpdateComment)
		commentRouter.DELETE("/:id", inDB.DeleteComment)
	}
	SocialMediaRouter := router.Group("/socialmedia")
	{
		SocialMediaRouter.Use(middlewares.Authentication())
		SocialMediaRouter.POST("/", inDB.StoreSocialMedia)
		SocialMediaRouter.GET("/", inDB.IndexSocialMedia)
		SocialMediaRouter.PUT("/:id", inDB.UpdateSocialMedia)
		SocialMediaRouter.DELETE("/:id", inDB.DeleteSocialMedia)
	}

	return router
}
