package config

import (
	"log"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

type Env struct {
	API_PORT          string
	API_LOG           string
	SQLDriver         string
	DATABASE          string
	DATABASE_USER     string
	DATABASE_PASSWORD string
	DATABASE_PORT     string
}

var Config Env

func init() {
	LoadEnv()
	Logging(Config.API_LOG)
}

func LoadEnv() {
	const projectDirName = "base-crud"
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkingDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkingDirectory))

	err := godotenv.Load(string(rootPath) + `/.env`)
	if err != nil {
		log.Fatalf("設定ファイルの読み込みに失敗しました。 %v", err)
	}
	Config = Env{
		API_PORT:          os.Getenv("API_PORT"),
		SQLDriver:         os.Getenv("SQL_DRIVER"),
		DATABASE:          os.Getenv("MYSQL_DATABASE"),
		DATABASE_USER:     os.Getenv("MYSQL_USER"),
		DATABASE_PASSWORD: os.Getenv("MYSQL_PASSWORD"),
		DATABASE_PORT:     os.Getenv("MYSQL_PORT"),
		API_LOG:           os.Getenv("API_LOG_FILE"),
	}
}
