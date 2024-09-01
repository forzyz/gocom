package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/lpernett/godotenv"
)

type Config struct {
	PublicHost            string
	Port                  string
	DBUser                string
	DBPassword            string
	DBAddress             string
	DBName                string
	JWTSecret             string
	JWTExpirationInSeconds int64
}

var Envs = initConfig()

func initConfig() Config {
	// Load environment variables from the .env file
	godotenv.Load()

	return Config{
		// Use environment variables without fallback values
		PublicHost: getEnv("PUBLIC_HOST", ""),
		Port:       getEnv("PORT", ""),
		DBUser:     getEnv("DB_USER", ""),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBAddress:  fmt.Sprintf("%s:%s", getEnv("DB_HOST", ""), getEnv("DB_PORT", "")),
		DBName:     getEnv("DB_NAME", ""),
		JWTSecret:  getEnv("JWT_SECRET", ""),
		JWTExpirationInSeconds: getEnvAsInt("JWT_EXP", 0),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}
		return i
	}
	return fallback
}