package config

import (
	"alba054/kartjis-notify/shared"
	"log"
	"os"
	"strconv"
	"strings"
)

type EnvConfig struct {
	DatabaseUrl string
	Host        string
	Port        int
}

func LoadConfig() EnvConfig {
	content, err := os.ReadFile(".env")
	shared.ThrowError(err)

	splittedContent := strings.Split(string(content), "\n")

	configMap := splitAndMapcontent(splittedContent, '=')

	var envConfig EnvConfig
	envConfig.DatabaseUrl = configMap["DATABASE_URL"]
	envConfig.Host = configMap["HOST"]
	port, _ := strconv.Atoi(configMap["PORT"])
	envConfig.Port = port

	return envConfig
}

func splitAndMapcontent(contents []string, sep rune) map[string]string {
	configMap := make(map[string]string)

	for _, v := range contents {
		split := strings.Split(v, string(sep))
		if len(split) < 2 {
			log.Fatal("cannot map a content")
		}
		configMap[split[0]] = split[1]
	}

	return configMap
}
