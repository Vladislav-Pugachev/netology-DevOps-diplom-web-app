package aws
import (
	"context"
	"fmt"
	"log"
	"time"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"os"
)

type Link string
func (l *Link) ReturnLink(path string) string{
	// Создаем кастомный обработчик эндпоинтов, который для сервиса S3 и региона ru-central1 выдаст корректный URL
	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		if service == s3.ServiceID && region == "ru-central1" {
			return aws.Endpoint{
				PartitionID:   "yc",
				URL:           "https://storage.yandexcloud.net",
				SigningRegion: "ru-central1",
			}, nil
		}
		return aws.Endpoint{}, fmt.Errorf("unknown endpoint requested")
	})
	cred:=config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
		Value: aws.Credentials{
			AccessKeyID: os.Getenv("AccessKeyID"), SecretAccessKey: os.Getenv("SecretAccessKey"),
		}})
	// Подгружаем конфигрурацию
	cfg, err := config.LoadDefaultConfig(context.TODO(), cred,config.WithEndpointResolverWithOptions(customResolver))
	if err != nil {
		log.Fatal(err)
	}

	// Создаем клиента для доступа к хранилищу S3
	client := s3.NewFromConfig(cfg)
	// Запрашиваем список бакетов

	presignClient := s3.NewPresignClient(client)
	presignResult, err := presignClient.PresignGetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String("pvv-netology-diplom-web"),
		Key:    aws.String(path),
	}, func(opts *s3.PresignOptions) {
		opts.Expires = time.Duration(30 * int64(time.Second))
	})
	if err != nil {
		panic("Невозможно создать ссылку на объект")
	}
	log.Printf("запрос на получение ссылки для картинки %s отправлен", path+".jpg")
	log.Printf("получена ссылка %s на картинку", presignResult.URL)
	return presignResult.URL

}
