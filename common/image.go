package common

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type Image struct {
	Id        int    `json:"id" gorm:"column:id;"`
	Url       string `json:"url" gorm:"column:url;"`
	Width     int    `json:"width" gorm:"column:width;"`
	Height    int    `json:"height" gorm:"column:height;"`
	CloudName string `json:"cloud_name, omitempty" gorm:"-"`
	Extention string `json:"extension, omitempty" gorm:"-"`
}

func (Image) TableName() string {
	return "image"
}

func (i *Image) Fulfill(domain string) {
	i.Url = fmt.Sprintf("%s/%s", domain, i.Url)
}

func (i *Image) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		errors.New(fmt.Sprint("Fail to unmarshal JSONB value: ", value))
	}

	var img Image
	if err := json.Unmarshal(bytes, &img); err != nil {
		return err
	}

	*i = img
	return nil
}

func (i *Image) Value() (driver.Value, error) {
	if i == nil {
		return nil, nil
	}
	return json.Marshal(i)
}

type Images = []Image

// func (i *Images) Scan(value interface{}) error {
// 	bytes, ok := value.([]byte)
// 	if !ok {
// 		errors.New(fmt.Sprint("Fail to unmarshal JSONB value: ", value))
// 	}

// 	var img []Image
// 	if err := json.Unmarshal(bytes, &img); err != nil {
// 		return err
// 	}

// 	*i = img
// 	return nil
// }

// func (i *Images) Value() (driver.Value, error) {
// 	if i == nil {
// 		return nil, nil
// 	}
// 	return json.Marshal(i)
// }
