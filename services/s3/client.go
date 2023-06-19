package s3

import (
	"github.com/Paxx-RnD/go-helper/concurrent"
	"github.com/minio/minio-go"
	"net/url"
	"time"
)

type IService interface {
	PreSignedGetObject(input string, credentials Credentials) (u *url.URL, err error)
	GetS3Client(endpoint string, accessKey string, secretKey string) (*minio.Client, error)
}

type service struct {
	clientMap     *concurrent.Dictionary[string, client]
	preSignedUrls *concurrent.Dictionary[string, *url.URL]
}

type client struct {
	minioClient *minio.Client
	lastUse     time.Time
}

func NewService() IService {
	s := service{
		clientMap:     concurrent.NewDictionary[string, client](),
		preSignedUrls: concurrent.NewDictionary[string, *url.URL](),
	}
	s.initClientS3GarbageCollector()
	return &s
}

func (s *service) initClientS3GarbageCollector() {
	go func() {
		for range time.Tick(time.Hour) {
			for _, key := range s.clientMap.Keys() {
				value, ok := s.clientMap.Get(key)
				if !ok {
					continue
				}
				if time.Since(value.lastUse) > time.Hour {
					s.clientMap.Delete(key)
				}
			}
		}
	}()
}

func (s *service) PreSignedGetObject(input string, credentials Credentials) (u *url.URL, err error) {
	if value, ok := s.preSignedUrls.Get(input); ok {
		return value, nil
	}

	s3Client, err := s.GetS3Client(credentials.Host, credentials.AccessKey, credentials.SecretKey)
	if err != nil {
		return nil, err
	}

	preSignedURL, err := s3Client.PresignedGetObject(credentials.Bucket, input, time.Hour, nil)
	if err != nil {
		return nil, err
	}

	s.preSignedUrls.Set(input, preSignedURL)
	go func() {
		time.Sleep(time.Second * 30)
		s.preSignedUrls.Delete(input)
	}()

	return preSignedURL, nil
}

func (s *service) GetS3Client(endpoint string, accessKey string, secretKey string) (*minio.Client, error) {
	var s3Client client
	var ok bool
	url, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	if s3Client, ok = s.clientMap.Get(endpoint); !ok {
		newClient, err := minio.New(url.Hostname(), accessKey, secretKey, false)
		if err != nil {
			return nil, err
		}
		s3Client = client{minioClient: newClient, lastUse: time.Now()}
		s.clientMap.Set(endpoint, s3Client)
	} else {
		s3Client.lastUse = time.Now()
	}
	return s3Client.minioClient, nil
}
