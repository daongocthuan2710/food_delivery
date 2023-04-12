package restaurantBiz

import (
	"context"
	"errors"
	"food_delivery/common"
	restaurantModel "food_delivery/modules/restaurant/model"
)

type DeleteRestaurantStore interface {
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

type deleteNewRestaurantBiz struct {
	store UpdateRestaurantStore
}

func NewDeleteRestaurantBiz(store UpdateRestaurantStore) *deleteNewRestaurantBiz {
	return &deleteNewRestaurantBiz{store: store}
}

func (biz *deleteNewRestaurantBiz) DeleteRestaurant(ctx context.Context, id int) error {

	olData, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		if err == common.ErrDataNotFound {
			return errors.New("data not found")
		}
		return common.ErrCannotGetEntity(restaurantModel.EntityName, err)
	}

	if olData.Status == 0 {
		return common.ErrEntityDeleted(restaurantModel.EntityName, err)
	}

	zero := 0

	if err := biz.store.Update(ctx,
		map[string]interface{}{"id": id},
		&restaurantModel.RestaurantUpdate{Status: &zero}); err != nil {
		return common.ErrCannotDeleteEntity(restaurantModel.EntityName, err)
	}

	return nil
}
