package main

import (
	"food_delivery/common/component/appctx"
	"food_delivery/common/component/uploadprovider"
	"food_delivery/middleware"
	"log"

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
	secretKey := os.Getenv("SYSTEM_SECRET")

	dns := os.Getenv("DB_CONN_STR")
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	//refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for detail7

	if err != nil {
		panic(err)
	}
	log.Println(db)
	db = db.Debug()

	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)
	appCtx := appctx.NewAppContext(db, s3Provider, secretKey)

	r := gin.Default()
	r.Use(middleware.Recover(appCtx))

	r.Static("/static/", "./static")
	mainRoute(r, appCtx)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
