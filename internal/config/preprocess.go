package config

import (
	"flag"
	"fmt"
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
	fmt.Printf("all before load:\n %#v\n", os.Environ())
	godotenv.Load()
	fmt.Printf("all after load:\n %#v\n", os.Environ())
	log.Print("job_id: ", os.Getenv("$CI_JOB_ID"))
	log.Print("url: ", os.Getenv("$DB_URL"))
	log.Print("host: ", os.Getenv("$DB_HOST"))

	re := regexp.MustCompile(`\$\{(.+?)\}`)
	replacedContent := re.ReplaceAllStringFunc(content, func(s string) string {
		envVarName := strings.TrimSuffix(strings.TrimPrefix(s, `${`), `}`)
		log.Print("try to get: ", envVarName, " ", os.Getenv(envVarName))
		envVarValue := os.Getenv(envVarName)
		return envVarValue
	})

	return []byte(replacedContent), err
}
