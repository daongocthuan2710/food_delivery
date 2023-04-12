package restaurantBiz

import (
	"context"
	"errors"
	"food_delivery/common"
	restaurantModel "food_delivery/modules/restaurant/model"
)

type UpdateRestaurantStore interface {
	FindDataWithCondition(
		ctx context.Context,
		cond map[string]interface{},
		moreKeys ...string,
	) (*restaurantModel.Restaurant, error)

	Update(
		ctx context.Context,
		cond map[string]interface{},
		updateData *restaurantModel.RestaurantUpdate,
	) error
}

type updateNewRestaurantBiz struct {
	store UpdateRestaurantStore
}

func NewUpdateRestaurantBiz(store UpdateRestaurantStore) *updateNewRestaurantBiz {
	return &updateNewRestaurantBiz{store: store}
}

func (biz *updateNewRestaurantBiz) UpdateRestaurant(ctx context.Context, id int, data *restaurantModel.RestaurantUpdate) error {
	if err := data.Validate(); err != nil {
		return common.ErrInvalidRequest(err)
	}

	olData, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		if err == common.ErrDataNotFound {
			return errors.New("data not found")
		}
		return common.ErrCannotGetEntity(restaurantModel.EntityName, err)
	}

	if olData.Status == 0 {
		return errors.New("Data has been deleted")
	}

	if err := biz.store.Update(ctx, map[string]interface{}{"id": id}, data); err != nil {
		return common.ErrCannotUpdateEntity(restaurantModel.EntityName, err)
	}

	return nil
}
