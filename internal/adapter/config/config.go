package config

type WasabiConfig struct {
	BucketName string
	AccessKey  string
	SecretKey  string
	Region     string
}

type HTTPConfig struct {
	AdminKey string
	Port     int
}
