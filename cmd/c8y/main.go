package main

import (
	"os"
	"time"

	"github.com/iqnev/golang-c8y-mqtt/internal"
	"github.com/iqnev/golang-c8y-mqtt/internal/common"
)

var exit = make(chan bool)

func main() {
	var env string
	if len(os.Args) == 2 {
		env = os.Args[1]
	}
	configuration := common.GetConfiguration(env)

	common.InitClientOptions(configuration)

	internal.RunDevice(configuration)

	for {
		//keep alive!!
		time.Sleep(2)
	}
}
