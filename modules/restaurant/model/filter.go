package restaurantModel

type Filter struct {
	UserId int `json:"-" form:"user_id"` // query string
}
