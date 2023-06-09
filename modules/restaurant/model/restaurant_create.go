package restaurantModel

import (
	// "food_delivery/common"
	"food_delivery/common"
	"strings"
)

type RestaurantCreate struct {
	common.SQLModel
	Name    string        `json:"name" gorm:"column:name;"`
	Address string        `json:"address" gorm:"column:addr;"`
	OwnerId int           `json:"-" gorm:"column:owner_id"`
	Logo    *common.Image `json:"logo" gorm:"column:logo;"`
	// Cover   *common.Images `json:"cover" gorm:"column:cover;"`
}

func (RestaurantCreate) TableName() string {
	return "restaurants"
}

func (data *RestaurantCreate) Validate() error {
	data.Name = strings.TrimSpace(data.Name)
	if data.Name == "" {
		return ErrNameCanNotBeEmpty
	}

	data.Address = strings.TrimSpace(data.Address)
	if data.Address == "" {
		return ErrAddressCanNotBeEmpty
	}

	return nil
}
