package docker_test

import (
	"os"
)

func InitEnv() {
	os.Setenv("HTTP_HOST", "0.0.0.0")
	os.Setenv("HTTP_PORT", "8888")
	os.Setenv("JWT_SECRET", "gFd4OE87FMVgd1BYcTVXVG")
	os.Setenv("JWT_TTL_SEC", "300")
	os.Setenv("PASSWORD_ITERATIONS", "3")
	os.Setenv("PASSWORD_KEY_LENGTH", "32")
	os.Setenv("PASSWORD_MEMORY", "65536")
	os.Setenv("PASSWORD_PARALLELISM", "2")
	os.Setenv("PASSWORD_SALT_LENGTH", "16")
	os.Setenv("PASSWORD_SECRET", "1mf5jeFSD%$Luf85[js")
	os.Setenv("POSTGRES_DB", "chatty")
	os.Setenv("POSTGRES_HOST", "0.0.0.0")
	os.Setenv("POSTGRES_PASSWORD", "postgres")
	os.Setenv("POSTGRES_POOL_MAX", "5")
	os.Setenv("POSTGRES_PORT", "5432")
	os.Setenv("POSTGRES_USER", "postgres")
}
