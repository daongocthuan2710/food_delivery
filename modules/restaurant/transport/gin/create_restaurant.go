package restaurantGin

import (
	"food_delivery/common/component/appctx"
	restaurantBiz "food_delivery/modules/restaurant/biz"
	restaurantModel "food_delivery/modules/restaurant/model"
	restaurantStorage "food_delivery/modules/restaurant/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateRestaurant(appCtx appctx.AppContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var newData restaurantModel.RestaurantCreate
		//&newData để xác định vị trí và tăng ID
		if err := ctx.ShouldBind(&newData); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// Dependencies install
		store := restaurantStorage.NewSqlStore(appCtx.GetMainDBConnection())
		biz := restaurantBiz.NewCreateNewRestaurantBiz(store)

		if err := biz.CreateNewRestaurant(ctx.Request.Context(), &newData); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": newData.Id})
	}
}
