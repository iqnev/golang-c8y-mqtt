package common

import (
	"fmt"
	"log"
	"runtime"

	"github.com/tkanos/gonfig"
)

type Configuration struct {
	C8Y_SEREVR_URL    string
	DEVICE_NAME       string
	C8Y_TENENT        string
	C8Y_USERNAME      string
	C8Y_PASSWORD      string
	CLIENT_ID         string
	DEVICE_OPERATIONS string
}

func GetConfiguration(params ...string) Configuration {
	configuration := Configuration{}
	env := "dev"

	if len(params) > 0 {
		env = params[0]
	}

	var configPath string

	os := runtime.GOOS
	switch os {
	case "windows":
		configPath = "C:/configs/"
	case "darwin":
		fmt.Println("MAC operating system")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s.\n", os)
	}

	if len(configPath) == 0 {
		log.Fatal("The configPath is empty!")
	}

	fileName := fmt.Sprintf(configPath+"%s_config.json", env)

	gonfig.GetConf(fileName, &configuration)

	return configuration
}
