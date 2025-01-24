package config

import (
	"flag"
	"github.com/joho/godotenv"
	"log"
	"os"
	"regexp"
	"strings"
)

func preprocess() ([]byte, error) {
	var (
		data []byte
		err  error
	)

	configFile := flag.String("config", "./config.yaml", "Path to config file")
	flag.Parse()

	data, err = os.ReadFile(*configFile)
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}

	content := string(data)
	godotenv.Load()
	log.Print("job_id: ", os.Getenv("CI_JOB_ID"))
	re := regexp.MustCompile(`\$\{(.+?)\}`)
	replacedContent := re.ReplaceAllStringFunc(content, func(s string) string {
		envVarName := strings.TrimSuffix(strings.TrimPrefix(s, `${`), `}`)
		log.Print("try to get: ", envVarName)
		envVarValue := os.Getenv(envVarName)
		return envVarValue
	})

	return []byte(replacedContent), err
}
