package restaurantBiz

import (
	"context"
	"food_delivery/common"
	restaurantModel "food_delivery/modules/restaurant/model"
)

type GetRestaurantStore interface {
	FindDataWithCondition(
		ctx context.Context,
		cond map[string]interface{},
		moreKeys ...string,
	) (*restaurantModel.Restaurant, error)
}

type getNewRestaurantBiz struct {
	store GetRestaurantStore
}

func NewGetRestaurantBiz(store GetRestaurantStore) *getNewRestaurantBiz {
	return &getNewRestaurantBiz{store: store}
}

func (biz *getNewRestaurantBiz) GetRestaurant(ctx context.Context, id int) (*restaurantModel.Restaurant, error) {

	result, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return nil, common.ErrCannotGetEntity(restaurantModel.EntityName, err)
	}

	return result, nil
}
