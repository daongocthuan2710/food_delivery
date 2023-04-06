package restaurantBiz

import (
	"context"
	"errors"
	restaurantModel "food_delivery/modules/restaurant/model"
	"strings"
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
	data.Name = strings.TrimSpace(data.Name)
	if data.Name == "" {
		return errors.New("restaurant name cannot be empty")
	}

	data.Address = strings.TrimSpace(data.Address)
	if data.Address == "" {
		return errors.New("restaurant address cannot be empty")
	}

	if err := biz.store.Create(ctx, data); err != nil {
		return err
	}

	return nil
}
