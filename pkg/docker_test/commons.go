package docker_test

import (
	"os"
)

func InitEnv() {
	os.Setenv("HTTP_PORT", "8888")
	os.Setenv("HTTP_HOST", "localhost")
	os.Setenv("POSTGRES_HOST", "localhost")
	os.Setenv("POSTGRES_PORT", "5432")
	os.Setenv("POSTGRES_USER", "postgres")
	os.Setenv("POSTGRES_PASSWORD", "postgres")
	os.Setenv("POSTGRES_DB", "chatty")
	os.Setenv("POSTGRES_POOL_MAX", "5")
	os.Setenv("JWT_SIGN", "secureJWTsign")
	os.Setenv("JWT_EXPIRED_MINUTES", "5")
}
