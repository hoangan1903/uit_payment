package env

import (
	"fmt"
	"os"
	"strconv"
)

func getEnv(env string) string {
	if value := os.Getenv(env); value == "" {
		panic(fmt.Sprintf("ENV %s is empty", env))
	} else {
		return value
	}
}

func GetDBHost() string {
	return getEnv("DB_HOST")
}

func GetDBPort() string {
	return getEnv("DB_PORT")
}

func GetDBSSLMode() string {
	return getEnv("DB_SSL_MODE")
}

func GetDBUserName() string {
	return getEnv("DB_USERNAME")
}

func GetDBPassword() string {
	return getEnv("DB_PASSWORD")
}

func GetDBName() string {
	return getEnv("DB_NAME")
}

func GetMaxOpenConns() int {
	if v, err := strconv.Atoi(getEnv("DB_MAX_OPEN_CONNS")); err == nil {
		return v
	}

	return 0
}

func GetMaxIdleConns() int {
	if v, err := strconv.Atoi(getEnv("DB_MAX_IDLE_CONNS")); err == nil {
		return v
	}

	return 0
}

func GetLogAge() string {
	return getEnv("LOG_AGE")
}

// func GetGraylogURL() string {
// 	return getEnv("GRAYLOG_ADDRESS")
// }

// func GetEnvironment() string {
// 	return getEnv("ENVIRONMENT")
// }

// func GetFacility() string {
// 	return getEnv("FACILITY")
// }

func GetLogLevel() string {
	return getEnv("LOG_LEVEL")
}

func GetFileMaxSize() int {
	v, err := strconv.Atoi(os.Getenv("FILE_MAX_SIZE"))
	if err != nil {
		return 0
	}
	return v
}

func GetFileBackups() int {
	v, err := strconv.Atoi(os.Getenv("FILE_BACKUPS"))
	if err != nil {
		return 0
	}
	return v
}

func GetFileMaxAge() int {
	v, err := strconv.Atoi(os.Getenv("FILE_MAX_AGE"))
	if err != nil {
		return 0
	}
	return v
}

func GetMomoCallbackURL() string {
	return getEnv("MOMO_CALLBACK_URL")
}

func GetMomoPartnerCode() string {
	return getEnv("MOMO_PARTNERCODE")
}

func GetMomoAccessKey() string {
	return getEnv("MOMO_ACCESSKEY")
}

func GetMomoSecretKey() string {
	return getEnv("MOMO_SECRETKEY")
}

func GetMomoAIOURL() string {
	return getEnv("MOMO_AIO_URL")
}

func GetMomoVersion() float32 {
	version := getEnv("MOMO_VERSION")
	value, err := strconv.ParseFloat(version, 32)
	if err != nil {
		// do something sensible
	}
	float := float32(value)
	return float
}
