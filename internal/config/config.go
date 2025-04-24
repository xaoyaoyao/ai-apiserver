/**
 * Package repo
 * @file      : config.go
 * @author    : xaoyaoyao
 * @contact   : xaoyaoyao@aliyun.com
 * @version   : 1.0.0
 * @time      : 2025/4/22 15:56
 **/

package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var cfg Config

type Config struct {
	Addr     string // 服务地址IP
	LogLevel string // 日志级别

	AccessKeyId     string
	SecretAccessKey string

	VolcEngineAddr    string
	VolcEnginePath    string
	VolcEngineService string
	VolcEngineRegion  string
	VolcEngineVersion string

	MeituApiKey      string
	MeituSecretKey   string
	MeituSyncPushUrl string
}

func Get() Config {
	return cfg
}

func Init() error {
	err := InitBaseConfig()
	if err != nil {
		fmt.Println("Error initializing base config:", err)
		return err
	}
	return nil
}

func InitBaseConfig() error {
	filePath := ".env"
	err := godotenv.Load(filePath)
	if err != nil {
		log.Println(err)
		log.Fatal("Error loading .env file")
	}
	return loadBaseConfigByEnv()
}

func loadBaseConfigByEnv() error {
	cfg = Config{
		Addr:              os.Getenv("ADDR"),
		LogLevel:          os.Getenv("LOG_LEVEL"),
		AccessKeyId:       os.Getenv("ACCESS_KEY_ID"),
		SecretAccessKey:   os.Getenv("SECRET_ACCESS_KEY"),
		VolcEngineAddr:    os.Getenv("VOLC_ENGINE_ADDR"),
		VolcEnginePath:    os.Getenv("VOLC_ENGINE_PATH"),
		VolcEngineService: os.Getenv("VOLC_ENGINE_SERVICE"),
		VolcEngineRegion:  os.Getenv("VOLC_ENGINE_REGION"),
		VolcEngineVersion: os.Getenv("VOLC_ENGINE_VERSION"),

		MeituApiKey:      os.Getenv("MEITU_API_KEY"),
		MeituSecretKey:   os.Getenv("MEITU_SECRET_KEY"),
		MeituSyncPushUrl: os.Getenv("MEITU_SYNC_PUSH_URL"),
	}
	return nil
}
