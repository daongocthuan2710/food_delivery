package restaurantModel

import (
	"food_delivery/common"
	"strings"
)

type RestaurantUpdate struct {
	Name    *string        `json:"name" gorm:"column:name;"`
	Address *string        `json:"address" gorm:"column:addr;"`
	Status  *int           `json:"-" gorm:"column:status;"`
	Cover   *common.Images `json:"cover" gorm:"column:cover;"`
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}

func (data *RestaurantUpdate) Validate() error {
	if strPtr := data.Name; strPtr != nil {
		str := strings.TrimSpace(*strPtr)
		if str == "" {
			return ErrNameCanNotBeEmpty
		}
		data.Name = &str
	}

	if strPtr := data.Address; strPtr != nil {
		str := strings.TrimSpace(*strPtr)
		if str == "" {
			return ErrAddressCanNotBeEmpty
		}
		data.Address = &str
	}

	return nil
}
