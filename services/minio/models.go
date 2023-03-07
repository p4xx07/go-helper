package minio

type Credentials struct {
	Id         int    `json:"id"`
	CustomerId int    `json:"customer_id"`
	S3Name     string `json:"s3_name"`
	Host       string `json:"host"`
	Bucket     string `json:"bucket"`
	AccessKey  string `json:"access_key"`
	SecretKey  string `json:"secret_key"`
}
