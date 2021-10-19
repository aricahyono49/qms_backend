package config

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/subosito/gotenv"
)

var (
	Configuration Config
)

type Config struct {
	AppName        string
	AppPort        string
	LogLevel       string
	Environment    string
	JWTSecret      string
	RedisAddress   string
	DBUsername     string
	DBPassword     string
	DBHost         string
	DBPort         int
	DBName         string
	MinioEndpoint  string
	MinioAccessKey string
	MinioSecretKey string
	MinioRegion    string
	MinioBucket    string
}

func Init() (config Config) {
	defaultEnv := ".env"

	if err := gotenv.Load(defaultEnv); err != nil {
		log.Warning("failed load .env")
	}

	log.SetOutput(os.Stdout)

	Configuration.AppName = GetString("APP_NAME")
	Configuration.AppPort = GetString("APP_PORT")
	Configuration.LogLevel = GetString("LOG_LEVEL")
	Configuration.Environment = GetString("ENVIRONMENT")
	Configuration.JWTSecret = GetString("JWT_SECRET")
	Configuration.RedisAddress = GetString("REDIS_ADDRESS")
	Configuration.DBUsername = GetString("DB_USERNAME")
	Configuration.DBPassword = GetString("DB_PASSWORD")
	Configuration.DBHost = GetString("DB_HOST")
	Configuration.DBPort = GetInt("DB_PORT")
	Configuration.DBName = GetString("DB_NAME")
	Configuration.MinioEndpoint = GetString("MINIO_ENDPOINT")
	Configuration.MinioAccessKey = GetString("MINIO_ACCESS_KEY")
	Configuration.MinioSecretKey = GetString("MINIO_SECRET_KEY")
	Configuration.MinioRegion = GetString("MINIO_REGION")
	Configuration.MinioBucket = GetString("MINIO_BUCKET")

	return Configuration
}
