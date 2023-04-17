package common

import (
	"errors"
	"fmt"
)

var ErrDataNotFound = errors.New("data not found")

var CurrentUser = "user"

const (
	DbTypeRestaurant = 1
	DbTypeUser       = 2
)

func Recovery() {
	if r := recover(); r != nil {
		fmt.Println("Recovered: ", r)
	}
}

type Requester interface {
	GetUserId() int
	GetEmail() string
	GetRole() string
}
