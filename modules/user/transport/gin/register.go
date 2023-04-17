package ginuser

import (
	"food_delivery/common"
	"food_delivery/common/component/appctx"
	hasher "food_delivery/common/component/hasher"
	userbiz "food_delivery/modules/user/biz"
	usermodel "food_delivery/modules/user/model"
	userstorage "food_delivery/modules/user/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(appCtx appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		var data usermodel.UserCreate

		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		store := userstorage.NewSqlStore(db)
		md5 := hasher.NewMd5Hash()
		biz := userbiz.NewRegisterBiz(store, md5)

		if err := biz.Register(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask(common.DbTypeUser)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
