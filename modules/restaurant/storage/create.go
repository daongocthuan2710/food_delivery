package restaurantStorage

import (
	"context"
	"food_delivery/common"
	restaurantModel "food_delivery/modules/restaurant/model"
)

func (s *sqlStore) Create(ctx context.Context, data *restaurantModel.RestaurantCreate) error {
	db := s.db
	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
