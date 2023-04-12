package restaurantBiz

import (
	"context"
	"food_delivery/common"
	restaurantModel "food_delivery/modules/restaurant/model"
)

type ListRestaurantStore interface {
	ListDataWithCondition(
		ctx context.Context,
		filter *restaurantModel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]restaurantModel.Restaurant, error)
}

type listNewRestaurantBiz struct {
	store ListRestaurantStore
}

func NewListRestaurantBiz(store ListRestaurantStore) *listNewRestaurantBiz {
	return &listNewRestaurantBiz{store: store}
}

func (biz *listNewRestaurantBiz) ListRestaurant(
	ctx context.Context,
	filter *restaurantModel.Filter,
	paging *common.Paging,
) ([]restaurantModel.Restaurant, error) {

	result, err := biz.store.ListDataWithCondition(ctx, filter, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantModel.EntityName, err)
	}

	return result, nil
}
