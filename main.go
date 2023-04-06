package main

import (
	"food_delivery/common"
	"food_delivery/common/component/appctx"
	restaurantGin "food_delivery/modules/restaurant/transport/gin"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Restaurant struct {
	common.SQLModel
	Name    string `json:"name" gorm:"column:name;"`
	Address string `json:"address" gorm:"column:addr;"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

type RestaurantCreate struct {
	common.SQLModel
	Name    string `json:"name" gorm:"column:name;"`
	Address string `json:"address" gorm:"column:addr;"`
}

func (RestaurantCreate) TableName() string {
	return "restaurants"
}

type RestaurantUpdate struct {
	common.SQLModel
	Name    *string `json:"name" gorm:"column:name;"`
	Address *string `json:"address" gorm:"column:addr;"`
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}

func main() {
	os.Setenv("DB_CONN_STR", "food_delivery:19e5a718a54a9fe0559dfbce6908@tcp(127.0.0.1:3307)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local")
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for detail7
	dsn := os.Getenv("DB_CONN_STR")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	log.Println(db)
	db = db.Debug()

	appCtx := appctx.NewAppContext(db)

	r := gin.Default()
	v1 := r.Group("/v1")
	{
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
