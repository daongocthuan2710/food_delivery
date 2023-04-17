package main

import (
	"food_delivery/common/component/appctx"
	"food_delivery/middleware"
	restaurantGin "food_delivery/modules/restaurant/transport/gin"
	ginupload "food_delivery/modules/upload/transport/gin"
	usergin "food_delivery/modules/user/transport/gin"
	userstorage "food_delivery/modules/user/storage"
	"github.com/gin-gonic/gin"
)

func mainRoute(r *gin.Engine, appCtx appctx.AppContext){
	authStore := userstorage.NewSqlStore(appCtx.GetMainDBConnection())
	v1 := r.Group("/v1")
	{
		v1.POST("/register", usergin.Register(appCtx))
		v1.POST("/login", usergin.Login(appCtx))
		v1.GET("/profile", middleware.RequiredAuth(appCtx, authStore), usergin.Profile(appCtx))

		v1.POST("/upload", ginupload.UploadImage(appCtx))
		// v1.GET("/presigned-upload-url", func(c *gin.Context) {
		// 	c.JSON(http.StatusOK, gin.H{"data": s3Provider.GetUpLoadPresignedUrl(c.Request.Context())})
		// }) // Client upload file image lên thẳng AWS có hiệu lực trong 60s

		restaurants := v1.Group("restaurants")
		{
			restaurants.POST("", restaurantGin.CreateRestaurant(appCtx))
			restaurants.GET("/:id", restaurantGin.GetRestaurant(appCtx))
			restaurants.PUT("/:id", restaurantGin.UpdateRestaurant(appCtx))
			restaurants.DELETE("/:id", restaurantGin.DeleteRestaurant(appCtx))
			restaurants.GET("", restaurantGin.ListRestaurant(appCtx))
		}
	}
}
