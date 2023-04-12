package restaurantGin

import (
	"food_delivery/common"
	"food_delivery/common/component/appctx"
	restaurantBiz "food_delivery/modules/restaurant/biz"
	restaurantModel "food_delivery/modules/restaurant/model"
	restaurantStorage "food_delivery/modules/restaurant/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListRestaurant(appCtx appctx.AppContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		go func() {
			defer common.Recovery()
			panic(1)
		}()

		var paging common.Paging
		if err := ctx.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var filter restaurantModel.Filter
		if err := ctx.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if err := paging.Process(); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := restaurantStorage.NewSqlStore(appCtx.GetMainDBConnection())
		biz := restaurantBiz.NewListRestaurantBiz(store)

		result, err := biz.ListRestaurant(ctx.Request.Context(), &filter, &paging)

		for i := range result {
			result[i].Mask(common.DbTypeRestaurant)
		}

		if err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
