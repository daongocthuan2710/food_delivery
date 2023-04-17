package userbiz

import (
	"context"
	"food_delivery/common"
	"food_delivery/common/component/tokenprovider"
	usermodel "food_delivery/modules/user/model"
)

type LoginStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

type loginBiz struct {
	storeUser     LoginStorage
	tokenProvider tokenprovider.Provider
	hasher        Hasher
	expiry        int
}

func NewLoginBiz(storeUser LoginStorage, tokenProvider tokenprovider.Provider, hasher Hasher, expiry int) *loginBiz {
	return &loginBiz{
		storeUser:     storeUser,
		tokenProvider: tokenProvider,
		hasher:        hasher,
		expiry:        expiry,
	}
}

// 1. Find user, email
// 2. Hash password from input and compare with password in db
// 3. Provider: issue JWT token for client
// 3.1. Access token and refresh token
// 4. Return token(s)

func (biz *loginBiz) Login(ctx context.Context, data *usermodel.UserLogin) (*tokenprovider.Token, error) {
	user, err := biz.storeUser.FindUser(ctx, map[string]interface{}{"email": data.Email})

	if err != nil {
		return nil, usermodel.ErrEmailOrPasswordInvalid
	}

	passHashed := biz.hasher.Hash(data.Password + user.Salt)

	if passHashed != user.Password {
		return nil, usermodel.ErrEmailOrPasswordInvalid
	}

	payload := tokenprovider.TokenPayload{
		UserId: user.Id,
		Role:   user.Role,
	}

	accessToken, err := biz.tokenProvider.Generate(payload, biz.expiry)

	if err != nil {
		return nil, common.ErrInternal(err)
	}

	// refreshToken, err := biz.tokenProvider.Generate(payload, biz.tkCfg.GetRtExp())
	// if err != nil {
	// 	return nil, common.ErrInternal(err)
	// }

	// account := usermodel.NewAccount(accessToken, refreshToken)

	return accessToken, nil
}
