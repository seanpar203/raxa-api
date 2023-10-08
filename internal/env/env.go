package env

import (
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var (
	// App config
	APP_ENV = GetEnv("APP_ENV", "dev")

	// Host
	SERVER_HOST = GetEnv("HOST", "localhost")
	SERVER_PORT = GetEnv("PORT", "8080")

	SSL_ENABLED = GetEnvAsBool("SSL_ENABLED", false)

	// Logging Config
	LOG_LEVEL = GetEnvAsInt("LOG_LEVEL", 1)

	// Media Config
	MEDIA_DIR      = getMediaDir(GetEnv("MEDIA_DIR", "media"))
	MEDIA_ENDPOINT = GetEnv("MEDIA_ENDPOINT", "/media/")

	// The amount of time it takes an access token to expire in Minutes
	ACCESS_TOKEN_EXPIRATION = time.Duration(GetEnvAsInt("ACCESS_TOKEN_EXPIRATION", 5)) * time.Minute

	// The amount of hours it takes for a refresh token to expire in Hours
	REFRESH_TOKEN_EXPIRATION = time.Duration(GetEnvAsInt("REFRESH_TOKEN_EXPIRATION", 168)) * time.Hour

	// PSQL
	PSQL_HOST              = GetEnv("PSQL_HOST", "localhost")
	PSQL_PORT              = GetEnvAsInt("PSQL_PORT", 5432)
	PSQL_USER              = GetEnv("PSQL_USER", "postgres")
	PSQL_PASS              = GetEnv("PSQL_PASS", "postgres")
	PSQL_DBNAME            = GetEnv("PSQL_DBNAME", "go_api")
	PSQL_SSQLMODE          = GetEnv("PSQL_SSLMODE", "disable")
	PSQL_MAX_OPEN_CONNS    = GetEnvAsInt("PSQL_MAX_OPEN_CONNS", 10)
	PSQL_MAX_IDLE_CONNS    = GetEnvAsInt("PSQL_MAX_IDLE_CONNS", 10)
	PSQL_CONN_MAX_LIFETIME = time.Duration(GetEnvAsInt("PSQL_CONN_MAX_LIFETIME", 60)) * time.Second
)

func getMediaDir(mediaDirName string) string {
	dir, _ := os.Getwd()
	return filepath.Join(dir, mediaDirName)
}

func GetEnv(key string, defaultVal string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}

	return defaultVal
}

func GetEnvAsInt(key string, defaultVal int) int {
	strVal := GetEnv(key, "")

	if val, err := strconv.Atoi(strVal); err == nil {
		return val
	}

	return defaultVal
}

func GetEnvAsBool(key string, defaultVal bool) bool {
	strVal := GetEnv(key, "")

	if val, err := strconv.ParseBool(strVal); err == nil {
		return val
	}

	return defaultVal
}

// GetEnvAsStringArr reads ENV and returns the values split by separator.
func GetEnvAsStringArr(key string, defaultVal []string, separator ...string) []string {
	strVal := GetEnv(key, "")

	if len(strVal) == 0 {
		return defaultVal
	}

	sep := ","
	if len(separator) >= 1 {
		sep = separator[0]
	}

	return strings.Split(strVal, sep)
}

// GetEnvAsStringArrTrimmed reads ENV and returns the whitespace trimmed values split by separator.
func GetEnvAsStringArrTrimmed(key string, defaultVal []string, separator ...string) []string {
	slc := GetEnvAsStringArr(key, defaultVal, separator...)

	for i := range slc {
		slc[i] = strings.TrimSpace(slc[i])
	}

	return slc
}

func GetEnvAsURL(key string, defaultVal string) *url.URL {
	strVal := GetEnv(key, "")

	if len(strVal) == 0 {
		u, err := url.Parse(defaultVal)
		if err != nil {
			panic("Failed to parse default value for env variable as URL")
		}

		return u
	}

	u, err := url.Parse(strVal)
	if err != nil {
		panic("Failed to parse env variable as URL")
	}

	return u
}
