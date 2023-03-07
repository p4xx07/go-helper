package minio

type Credentials struct {
	Host      string `json:"host"`
	Bucket    string `json:"bucket"`
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
}

func NewCredentials(host, bucket, accessKey, secretKey string) *Credentials {
	return &Credentials{
		Host:      host,
		Bucket:    bucket,
		AccessKey: accessKey,
		SecretKey: secretKey,
	}
}
