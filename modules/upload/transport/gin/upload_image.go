package ginupload

import (
	"fmt"
	"food_delivery/common/component/appctx"
	uploadbiz "food_delivery/modules/upload/biz"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadImage(appCtx appctx.AppContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		fileHeader, err := ctx.FormFile("file")
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		folder := ctx.DefaultPostForm("folder", "img")
		file, err := fileHeader.Open()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		defer file.Close() // defer close when function is done

		dataBytes := make([]byte, fileHeader.Size)
		if _, err := file.Read(dataBytes); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		biz := uploadbiz.NewUploadBiz(appCtx.UploadProvider())
		img, err := biz.Upload(ctx.Request.Context(), dataBytes, folder, fileHeader.Filename)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.SaveUploadedFile(fileHeader, fmt.Sprintf("./static/%s", fileHeader.Filename))
		ctx.JSON(http.StatusOK, gin.H{"data": img})
		// ctx.JSON(http.StatusOK, gin.H{"url": fmt.Sprintf("http://localhost:8080/static/%s", fileHeader.Filename)})
	}
}
