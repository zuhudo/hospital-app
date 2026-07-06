package config

import "os"

type Config struct {
	Port           string
	DBHost         string
	DBPort         string
	DBUser         string
	DBPassword     string
	DBName         string
	JWTSecret      string
	JWTExpiryHours int
	AllowedOrigins string
}

func Load() *Config {
	return &Config{
		Port:           getEnv("PORT", "3000"),
		DBHost:         getEnv("DB_HOST", "localhost"),
		DBPort:         getEnv("DB_PORT", "5432"),
		DBUser:         getEnv("DB_USER", "hospital"),
		DBPassword:     getEnv("DB_PASSWORD", "hospital_secret"),
		DBName:         getEnv("DB_NAME", "hospital_db"),
		JWTSecret:      getEnv("JWT_SECRET", "super-secret-key-change-in-production"),
		JWTExpiryHours: 72,
		AllowedOrigins: getEnv("ALLOWED_ORIGINS", "*"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
