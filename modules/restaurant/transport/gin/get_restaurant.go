package restaurantGin

import (
	"food_delivery/common/component/appctx"
	restaurantBiz "food_delivery/modules/restaurant/biz"
	restaurantStorage "food_delivery/modules/restaurant/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetRestaurant(appCtx appctx.AppContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		store := restaurantStorage.NewSqlStore(appCtx.GetMainDBConnection())
		biz := restaurantBiz.NewGetRestaurantBiz(store)

		data, err := biz.GetRestaurant(ctx.Request.Context(), id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": data})
	}
}
