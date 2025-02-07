package minio

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/Paxx-RnD/go-helper/concurrent"
	miniogo "github.com/minio/minio-go"
	"net/url"
	"os"
	"time"
)

type IService interface {
	PreSignedGetObject(input string, credentials Credentials) (u *url.URL, err error)
	GetS3Client(credentials Credentials) (*miniogo.Client, error)
	PutObject(file *os.File, destination string, credentials Credentials, userMetadata map[string]string) error
	PutBytesObject(data []byte, destination string, credentials Credentials, userMetadata map[string]string) error
	CopyObject(origin string, destination string, credentials Credentials, userMetadata map[string]string) error
	Exists(key string, credentials Credentials) (bool, error)
	DeleteObject(path string, credentials Credentials) error
	GetObject(path string, credentials Credentials) ([]byte, error)
}

type service struct {
	clientMap     *concurrent.Dictionary[string, clientWrapper]
	preSignedUrls *concurrent.Dictionary[string, *url.URL]
}

type clientWrapper struct {
	minioClient *miniogo.Client
	lastUse     time.Time
}

func NewService() IService {
	s := service{
		clientMap:     concurrent.NewDictionary[string, clientWrapper](),
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

	s3Client, err := s.GetS3Client(credentials)
	if err != nil {
		return nil, err
	}

	preSignedUrl, err := s3Client.PresignedGetObject(credentials.Bucket, input, time.Hour, nil)
	if err != nil {
		return nil, err
	}

	s.preSignedUrls.Set(input, preSignedUrl)
	go func() {
		time.Sleep(time.Second * 30)
		s.preSignedUrls.Delete(input)
	}()

	return preSignedUrl, nil
}

func (s *service) GetS3Client(credentials Credentials) (*miniogo.Client, error) {
	var s3Client clientWrapper
	var ok bool
	url, err := url.Parse(credentials.Host)
	if err != nil {
		return nil, err
	}

	key := fmt.Sprintf("%s_%s_%s_%s", credentials.Host, credentials.Bucket, credentials.AccessKey, credentials.SecretKey)
	if s3Client, ok = s.clientMap.Get(key); !ok {
		newClient, err := miniogo.New(url.Hostname(), credentials.AccessKey, credentials.SecretKey, false)
		if err != nil {
			return nil, err
		}
		s3Client = clientWrapper{minioClient: newClient, lastUse: time.Now()}
		s.clientMap.Set(credentials.Host, s3Client)
	} else {
		s3Client.lastUse = time.Now()
	}

	return s3Client.minioClient, nil
}

func (s *service) PutObject(file *os.File, destination string, credentials Credentials, userMetadata map[string]string) error {
	if file == nil {
		return errors.New("file is nil")
	}
	client, err := s.GetS3Client(credentials)
	if err != nil {
		return err
	}
	fileStat, err := file.Stat()
	if err != nil {
		return err
	}

	_, err = client.PutObject(
		credentials.Bucket,
		destination,
		file,
		fileStat.Size(),
		miniogo.PutObjectOptions{ContentType: "application/octet-stream", UserMetadata: userMetadata})

	if err != nil {
		return err
	}

	return nil
}

func (s *service) PutBytesObject(data []byte, destination string, credentials Credentials, userMetadata map[string]string) error {
	if data == nil {
		return errors.New("data is nil")
	}

	client, err := s.GetS3Client(credentials)
	if err != nil {
		return err
	}

	_, err = client.PutObject(
		credentials.Bucket,
		destination,
		bytes.NewReader(data),
		int64(len(data)),
		miniogo.PutObjectOptions{ContentType: "application/octet-stream", UserMetadata: userMetadata})

	if err != nil {
		return err
	}

	return nil
}

func (s *service) CopyObject(origin string, destination string, credentials Credentials, userMetadata map[string]string) error {
	client, err := s.GetS3Client(credentials)

	destInfo, err := miniogo.NewDestinationInfo(credentials.Bucket, destination, nil, userMetadata)
	if err != nil {
		return err
	}

	srcInfo := miniogo.NewSourceInfo(credentials.Bucket, origin, nil)
	err = client.CopyObject(destInfo, srcInfo)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) Exists(key string, credentials Credentials) (bool, error) {
	client, err := s.GetS3Client(credentials)
	if err != nil {
		return false, err
	}

	_, err = client.StatObject(credentials.Bucket, key, miniogo.StatObjectOptions{})
	if err != nil {
		if miniogo.ToErrorResponse(err).Code == "NoSuchKey" {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (s *service) DeleteObject(path string, credentials Credentials) error {
	s3Client, err := s.GetS3Client(credentials)
	if err != nil {
		return err
	}

	err = s3Client.RemoveObject(credentials.Bucket, path)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetObject(path string, credentials Credentials) ([]byte, error) {
	s3Client, err := s.GetS3Client(credentials)
	if err != nil {
		return nil, err
	}

	res, err := s3Client.GetObject(credentials.Bucket, path, miniogo.GetObjectOptions{})
	if err != nil {
		return nil, err
	}

	defer res.Close()

	var fileContents []byte
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(res)
	if err != nil {
		return nil, err
	}
	fileContents = buf.Bytes()

	return fileContents, nil
}
