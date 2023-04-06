package restaurantGin

import (
	"food_delivery/common/component/appctx"
	restaurantBiz "food_delivery/modules/restaurant/biz"
	restaurantModel "food_delivery/modules/restaurant/model"
	restaurantStorage "food_delivery/modules/restaurant/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateRestaurant(appCtx appctx.AppContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var data restaurantModel.RestaurantUpdate

		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		store := restaurantStorage.NewSqlStore(appCtx.GetMainDBConnection())
		biz := restaurantBiz.NewUpdateRestaurantBiz(store)

		if err := biz.UpdateRestaurant(ctx.Request.Context(),id, &data); err!= nil {
            ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

		ctx.JSON(http.StatusOK, gin.H{"data": "SUCCESS"})
	}
}
