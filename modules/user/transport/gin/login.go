package ginuser

import (
	"food_delivery/common"
	"food_delivery/common/component/appctx"
	"food_delivery/common/component/hasher"
	usermodel "food_delivery/modules/user/model"
	userbiz "food_delivery/modules/user/biz"
	userstorage "food_delivery/modules/user/storage"
	jwt "food_delivery/common/component/tokenprovider/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginUserData usermodel.UserLogin

		if err := c.ShouldBind(&loginUserData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()
		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())

		store := userstorage.NewSqlStore(db)
		md5 := hasher.NewMd5Hash()
		expiry := 60 * 60 * 24 * 30

		biz := userbiz.NewLoginBiz(store, tokenProvider, md5, expiry)
		account, err := biz.Login(c.Request.Context(), &loginUserData)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(account))
	}
}
