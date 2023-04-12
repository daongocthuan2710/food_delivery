package common

import (
	"errors"
	"fmt"
)

var ErrDataNotFound = errors.New("data not found")

const (
	DbTypeRestaurant = 1
)

func Recovery() {
	if r := recover(); r != nil {
		fmt.Println("Recovered: ", r)
	}
}
