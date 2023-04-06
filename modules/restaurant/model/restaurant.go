package restaurantModel

import (
	"errors"
	"food_delivery/common"
)

type Restaurant struct {
	common.SQLModel
	Name    string `json:"name" gorm:"column:name;"`
	Address string `json:"address" gorm:"column:addr;"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

var (
	ErrNameCanNotBeEmpty = errors.New("name can not be empty")
	ErrAddressCanNotBeEmpty = errors.New("address can not be empty") 
)
