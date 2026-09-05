package config

import (
	"fmt"
	"os"
)

type Config struct {
	AppPort              string
	DBHost               string
	DBPort               string
	DBUser               string
	DBPass               string
	DBName               string
	DBSSLMode            string
	JWTKey               string
	FileServiceUrl       string
	FileServiceSecretKey string
	VapidPublicKey       string
	VapidPrivateKey      string
	VapidSubject         string
}

func New() *Config {
	return &Config{
		AppPort:              getEnv("APP_PORT", "8080"),
		DBHost:               getEnv("DB_HOST", "localhost"),
		DBPort:               getEnv("DB_PORT", "5432"),
		DBUser:               getEnv("DB_USER", "dbuser"),
		DBPass:               getEnv("DB_PASSWORD", ""),
		DBName:               getEnv("DB_NAME", "dbname"),
		DBSSLMode:            getEnv("DB_SSLMODE", "disable"),
		JWTKey:               getEnv("JWT_SECRET", "secret-key"),
		FileServiceUrl:       getEnv("FILE_SERVICE_URL", "http://localhost:5000"),
		FileServiceSecretKey: getEnv("FILE_SERVICE_SECRET_KEY", "file-service-secret-key"),
		VapidPublicKey:       getEnv("VAPID_PUBLIC_KEY", "file-service-secret-key"),
		VapidPrivateKey:      getEnv("VAPID_PRIVATE_KEY", "file-service-secret-key"),
		VapidSubject:         getEnv("VAPID_SUBJECT", "file-service-secret-key"),
	}
}

func (c *Config) DBURL() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		c.DBUser, c.DBPass, c.DBHost, c.DBPort, c.DBName, c.DBSSLMode,
	)
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}
