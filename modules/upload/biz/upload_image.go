package uploadbiz

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"food_delivery/common"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"path/filepath"
	"strings"
	"time"
)

type uploadBiz struct {
	provider UploadProvider
}

type UploadProvider struct {
	SaveFileUploaded func(ctx context.Context, data []byte, dst string) (*common.Image, error)
}

func NewUploadBiz(provider UploadProvider) *uploadBiz {
	return &uploadBiz{provider: provider}
}

func (biz *uploadBiz) Upload(ctx context.Context, data []byte, folder, fileName string) (*common.Image, error) {
	fileBytes := bytes.NewReader(data)

	w, h, err := getImageDimesion(fileBytes)

	if err != nil {
		return nil, errors.New("file is not image")
	}

	if strings.TrimSpace(folder) == "" {
		folder = "img"
	}

	fileExt := filepath.Ext(fileName)
	fileName = fmt.Sprintf("%d%s", time.Now().UnixNano(), fileExt)

	img, err := biz.provider.SaveFileUploaded(ctx, data, fmt.Sprintf("%s/%s", folder, fileName))

	if err != nil {
		return nil, errors.New("cannot save image")
	}

	img.Width = w
	img.Height = h

	img.Extention = fileExt
	return img, nil
}

func getImageDimesion(reader *bytes.Reader) (int, int, error) {
	img, _, err := image.DecodeConfig(reader)
	if err != nil {
		log.Println("err: ", err)
		return 0, 0, err
	}
	return img.Width, img.Height, nil
}
