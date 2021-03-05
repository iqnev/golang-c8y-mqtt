package main

import (
	"time"

	"github.com/iqnev/golang-c8y-mqtt/internal"
	"github.com/iqnev/golang-c8y-mqtt/internal/common"
	log "github.com/sirupsen/logrus"
)

var exit = make(chan bool)

func main() {

	log.Info("Dev Configuration")

	configuration := common.GetConfiguration()

	common.InitClientOptions(configuration)

	println(configuration.C8Y_SEREVR_URL)

	internal.RunDevice(configuration)

	for {
		//keep alive!!
		time.Sleep(2)
	}
}
