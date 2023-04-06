package main

import (
	"food_delivery/common"
	"log"
	"net/http"
	"os"
	"strconv"

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
	// var oldRes Restaurant
	// //db.Table(oldRes.TableName()).Select("id, name").Where(map[string]interface{}{"id": 3}) // truy vấn nhanh hơn

	// emptyStr := ""
	// dataUpdate := RestaurantUpdate{
	// 	Name: &emptyStr, // cần một con trỏ tới string
	// }

	// db.Exec("select * from")

	r := gin.Default()
	v1 := r.Group("/v1")
	{
		restaurants := v1.Group("restaurants")
		{
			restaurants.POST("", func(ctx *gin.Context) {
				var newData RestaurantCreate
				//&newData để xác định vị trí và tăng ID
				if err := ctx.ShouldBind(&newData); err != nil {
					ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}

				if err := db.Create(&newData).Error; err != nil {
					ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					return
				}
				ctx.JSON(http.StatusOK, gin.H{"data": newData.Id})
			})

			restaurants.GET("/:id", func(ctx *gin.Context) {
				var data Restaurant

				id, err := strconv.Atoi(ctx.Param("id"))
				if err != nil {
					ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}

				if err := ctx.ShouldBind(&data); err != nil {
					ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}

				if err := db.Where("id = ?", id).First(&data).Error; err != nil {
					ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					return
				}
				ctx.JSON(http.StatusOK, gin.H{"data": data})
			})

			restaurants.PUT("/:id", func(ctx *gin.Context) {
				id, err := strconv.Atoi(ctx.Param("id"))
				if err != nil {
					ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}

				var data RestaurantUpdate

				if err := ctx.ShouldBind(&data); err != nil {
					ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}

				if err := db.Where("id = ?", id).Updates(&data).Error; err != nil {
					ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					return
				}
				ctx.JSON(http.StatusOK, gin.H{"data": "SUCCESS"})
			})

			restaurants.DELETE("/:id", func(ctx *gin.Context) {
				id, err := strconv.Atoi(ctx.Param("id"))

				if err != nil {
					ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}

				if err := db.Table(Restaurant{}.TableName()).Where("id = ?", id).Delete(nil).Error; err != nil {
					ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					return
				}
				ctx.JSON(http.StatusOK, gin.H{"data": "SUCCESS"})
			})

			restaurants.GET("", func(ctx *gin.Context) {
				var data []Restaurant

				if err := ctx.ShouldBind(&data); err != nil {
					ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}

				type Paging struct {
					Page  int `json:"page" format:"page"`
					Limit int `json:"limit" format:"limit"`
				}
				var paging Paging
				if err := ctx.ShouldBind(&paging); err != nil {
					ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}

				if paging.Page < 1 {
					paging.Page = 1
				}

				if paging.Limit < 1 {
					paging.Limit = 10
				}

				offset := (paging.Page - 1) * paging.Limit

				if err := db.Offset(offset).Limit(paging.Limit).Order("id desc").Find(&data).Error; err != nil {
					ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					return
				}
				ctx.JSON(http.StatusOK, gin.H{"data": data})
			})
		}
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
