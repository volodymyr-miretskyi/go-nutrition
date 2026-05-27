package config

import (
	"fmt"
	"os"
)

type Config struct {
	APP_ENV string
	Port    string
	DB      *db
	JWT     *jwt
	S3      *s3
	OpenAi  *openAi
}

type db struct {
	DbHost     string
	DbPort     string
	DbUser     string
	DbName     string
	DbPassword string
	Dsn        string
}

type jwt struct {
	Secret string
}

type s3 struct {
	Bucket string
	Region string
}

type openAi struct {
	ApiKey string
}

func initDb() *db {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	name := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")

	dsn := os.Getenv("DB_URL")

	fmt.Print("DSN ", dsn)
	return &db{DbHost: host, DbPort: port, DbUser: user, DbName: name, DbPassword: password, Dsn: dsn}
}

func initS3() *s3 {
	bucket := os.Getenv("S3_BUCKET")
	region := os.Getenv("S3_REGION")

	return &s3{
		Bucket: bucket,
		Region: region,
	}
}

func initJwt() *jwt {
	secret := os.Getenv("JWT_SECRET")

	return &jwt{
		Secret: secret,
	}
}

func initOpenAi() *openAi {
	apiKey := os.Getenv("OPENAI_API_KEY")

	return &openAi{
		ApiKey: apiKey,
	}
}

func LoadConfig() *Config {
	appEnv := os.Getenv("APP_ENV")
	port := os.Getenv("PORT")

	db := initDb()
	s3 := initS3()
	jwt := initJwt()
	openAI := initOpenAi()

	return &Config{
		APP_ENV: appEnv,
		Port:    port,
		DB:      db,
		S3:      s3,
		JWT:     jwt,
		OpenAi:  openAI,
	}
}
