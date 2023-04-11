package uploadprovider

import (
	"bytes"
	"context"
	"fmt"
	"food_delivery/common"
	"log"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type s3Provider struct {
	buckerName string
	region     string
	apiKey     string
	secret     string
	domain     string
	session    *session.Session
}

func NewS3Provider(buckerName string, region string, apiKey string, secret string, domain string) *s3Provider {
	provider := &s3Provider{
		buckerName: buckerName,
		region:     region,
		apiKey:     apiKey,
		secret:     secret,
		domain:     domain,
	}

	s3Session, err := session.NewSession(
		&aws.Config{
			Region: aws.String(provider.region),
			Credentials: credentials.NewStaticCredentials(
				provider.apiKey, // Asset key ID
				provider.secret, // Secret access key
				"",              // token can be ignored
			),
		})

	if err != nil {
		log.Fatalln(err)
	}

	provider.session = s3Session

	return provider

}

func (p *s3Provider) SaveFileUploaded(ctx context.Context, data []byte, dst string) (*common.Image, error) {
	fileBytes := bytes.NewReader(data)
	fileType := http.DetectContentType(data)

	_, err := s3.New(p.session).PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(p.buckerName),
		Key:         aws.String(dst),
		Body:        fileBytes,
		ACL:         aws.String("private"),
		ContentType: aws.String(fileType),
	})

	if err != nil {
		return nil, err
	}

	img := &common.Image{
		Url:       fmt.Sprintf("%s/%s", p.domain, dst),
		CloudName: "s3",
	}

	return img, nil
}

func (p *s3Provider) GetUpLoadPresignedUrl(ctx context.Context) string {
	req, _ := s3.New(p.session).PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String(p.buckerName),
		Key:    aws.String(fmt.Sprintf("img/&d", time.Now().UnixNano())),
		ACL:    aws.String("private"),
	})

	url, err := req.Presign(time.Second * 60)

	if err != nil {
		log.Println(err)
	}

	return url
}

func (p *s3Provider) GetDomain() string {
	return p.domain
}
