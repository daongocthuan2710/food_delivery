package main

import (
	"food_delivery/common/component/appctx"
	"food_delivery/common/component/uploadprovider"
	restaurantGin "food_delivery/modules/restaurant/transport/gin"
	ginupload "food_delivery/modules/upload/transport/gin"
	"log"
	"net/http"

	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading.env file")
	}

	s3BucketName := os.Getenv("S3_BUCKET_NAME")
	s3Region := os.Getenv("S3_REGION")
	s3APIKey := os.Getenv("S3_ACCESS_KEY")
	s3SecretKey := os.Getenv("S3_SECRET_KEY")
	s3Domain := os.Getenv("S3_DOMAIN")

	dns := os.Getenv("DB_CONN_STR")
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	//refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for detail7

	if err != nil {
		panic(err)
	}
	log.Println(db)
	db = db.Debug()

	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)
	appCtx := appctx.NewAppContext(db, s3Provider)

	r := gin.Default()
	r.Static("/static/", "./static")
	v1 := r.Group("/v1")
	{
		v1.POST("/upload", ginupload.UploadImage(appCtx))
		v1.GET("/presigned-upload-url", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"data": s3Provider.GetUpLoadPresignedUrl(c.Request.Context())})
		})
		// Client upload file image lên thẳng AWS có hiệu lực trong 60s
		restaurants := v1.Group("restaurants")
		{
			restaurants.POST("", restaurantGin.CreateRestaurant(appCtx))
			restaurants.GET("/:id", restaurantGin.GetRestaurant(appCtx))
			restaurants.PUT("/:id", restaurantGin.UpdateRestaurant(appCtx))
			restaurants.DELETE("/:id", restaurantGin.DeleteRestaurant(appCtx))
			restaurants.GET("", restaurantGin.ListRestaurant(appCtx))
		}
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
