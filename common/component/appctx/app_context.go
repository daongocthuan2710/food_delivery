package appctx

import (
	"food_delivery/common/component/uploadprovider"

	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	UploadProvider() uploadprovider.UploadProvider
}

type appContext struct {
	db             *gorm.DB
	uploadProvider uploadprovider.UploadProvider
}

func NewAppContext(db *gorm.DB) *appContext {
	return &appContext{db: db}
}

func (appCtx *appContext) GetMainDBConnection() *gorm.DB {
	return appCtx.db
}

func (appCtx *appContext) UploadProvider() uploadprovider.UploadProvider {
	return appCtx.uploadProvider
}
