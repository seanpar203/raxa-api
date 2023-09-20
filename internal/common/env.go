package common

import (
	"net/url"
	"os"
	"strconv"
	"strings"
)

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
			log.Panic().Str("key", key).Str("defaultVal", defaultVal).Err(err).Msg("Failed to parse default value for env variable as URL")
		}

		return u
	}

	u, err := url.Parse(strVal)
	if err != nil {
		log.Panic().Str("key", key).Str("strVal", strVal).Err(err).Msg("Failed to parse env variable as URL")
	}

	return u
}
