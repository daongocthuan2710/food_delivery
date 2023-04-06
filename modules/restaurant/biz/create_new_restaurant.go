package restaurantBiz

import (
	"context"
	restaurantModel "food_delivery/modules/restaurant/model"
)

type CreateRestaurantStore interface {
	Create(ctx context.Context, data *restaurantModel.RestaurantCreate) error
}

type createNewRestaurantBiz struct {
	store CreateRestaurantStore
}

func NewCreateNewRestaurantBiz(store CreateRestaurantStore) *createNewRestaurantBiz {
	return &createNewRestaurantBiz{store: store}
}

func (biz *createNewRestaurantBiz) CreateNewRestaurant(ctx context.Context, data *restaurantModel.RestaurantCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	if err := biz.store.Create(ctx, data); err != nil {
		return err
	}

	return nil
}
