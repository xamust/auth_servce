package config

import (
	"flag"
	"github.com/joho/godotenv"
	"os"
	"regexp"
	"strings"
)

func preprocess() ([]byte, error) {
	configFile := flag.String("config", "./config.yaml", "Path to config file")
	flag.Parse()
	data, err := os.ReadFile(*configFile)
	if err != nil {
		return nil, err
	}

	content := string(data)
	godotenv.Load()
	re := regexp.MustCompile(`\$\{(.+?)\}`)
	replacedContent := re.ReplaceAllStringFunc(content, func(s string) string {
		envVarName := strings.TrimSuffix(strings.TrimPrefix(s, `${`), `}`)
		envVarValue := os.Getenv(envVarName)
		return envVarValue
	})

	return []byte(replacedContent), err
}
