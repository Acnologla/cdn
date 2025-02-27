package storage

import (
	"context"
	"io"
	"strings"

	"github.com/Acnologla/cdn/internal/adapter/config"
	"github.com/Acnologla/cdn/internal/core/port"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type Wasabi struct {
	session *session.Session
	client  *s3.S3
	bucket  string
}

func (w *Wasabi) Upload(ctx context.Context, key string, seeker io.ReadSeeker) error {
	_, err := w.client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(w.bucket),
		Key:    aws.String(key),
		Body:   seeker,
	})
	return err
}

func (w *Wasabi) Get(ctx context.Context, key string) ([]byte, error) {
	downloader := s3manager.NewDownloader(w.session)
	buf := aws.NewWriteAtBuffer([]byte{})

	_, err := downloader.Download(buf,
		&s3.GetObjectInput{
			Bucket: aws.String(w.bucket),
			Key:    aws.String(key),
		})

	return buf.Bytes(), err
}

func NewWasabiStorage(ctx context.Context, config config.WasabiConfig) port.Storage {
	region := strings.Split(config.Endpoint, ".")[1]
	sess, err := session.NewSession(&aws.Config{
		Endpoint:    &config.Endpoint,
		Region:      &region,
		Credentials: credentials.NewStaticCredentials(config.AccessKey, config.SecretKey, ""),
	})
	if err != nil {
		panic(err)
	}
	svc := s3.New(sess)
	return &Wasabi{
		session: sess,
		client:  svc,
		bucket:  config.BucketName,
	}
}
