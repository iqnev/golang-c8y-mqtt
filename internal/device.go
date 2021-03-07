package internal

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/iqnev/golang-c8y-mqtt/internal/common"
	log "github.com/sirupsen/logrus"
)

const serialNumber string = "111-222-34"
const hardwareModel string = "Golang:123"
const reversion string = "1.1"
const requiredInterval string = "60"

var operations = []string{"c8y_Restart", "c8y_SoftwareList"}
var configuration common.Configuration

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	topic := msg.Topic()
	payload := msg.Payload()
	if strings.Compare(string(payload), "\n") > 0 {
		fmt.Printf("TOPIC: %s\n", topic)
		fmt.Printf("MSG: %s\n", payload)
	}

	if strings.Compare("bye\n", string(payload)) == 0 {
		fmt.Println("exitting")
		//	flag = true
	}
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

func RunDevice(conf common.Configuration) {
	log.Info("Starting device...")
	configuration = conf

	registerDevice()

	createSmartRestTemplates()

	hearDevice()

	go temperatureService()
}

func hearDevice() {
	topic := "s/ds"
	client := common.GetInstance().GetMqqtClient()
	if token := client.Subscribe(topic, 0, messagePubHandler); token.Wait() && token.Error() != nil {
		log.Error("An error occured while executing: %s", token.Error())
		common.CloseConnection()
		os.Exit(1)
	}
}

func registerDevice() {
	log.Info("Try to registering Device with serial: " + configuration.SERIAL_NUMBER)

	client := common.GetInstance().GetMqqtClient()
	client.Publish("s/us", 0, false, createDevice(configuration.DEVICE_NAME, "c8y_MQTTDevice"))

	time.Sleep(1 * time.Second)

	client.Publish("s/us", 0, false, createHardwareInfo(configuration.SERIAL_NUMBER, configuration.HARDWARE_MODEL, configuration.REVESION))

	time.Sleep(1 * time.Second)

	client.Publish("s/us", 0, false, createOperations(configuration.DEVICE_OPERATIONS))
}

func createSmartRestTemplates() {
	log.Info("Creating SmartResetTemplates.")
	//TODO
}

func sub(client mqtt.Client, topic string) {
	if token := client.Subscribe(topic, 0, messagePubHandler); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}
}

func createOperations(operations string) string {
	return "114," + operations
}
func createRequiredInterval(requiredInterval string) string {
	return "117," + requiredInterval
}

func createHardwareInfo(serialNumber string, hardwareModel string, reversion string) string {
	return "110," + serialNumber + "," + hardwareModel + "," + reversion
}

func createDevice(name string, deviceType string) string {
	return "100," + name + "," + deviceType
}

func createTemperatureMeasurement(value string, time string) string {
	msg := "211," + value
	if time != "" {
		msg = msg + "," + time
	}
	return msg
}

//Services
func temperatureService() {

	for {
		temp := strconv.Itoa(randInt(10, 33))
		log.Info("Sending temeperature meserment: " + temp + "C")
		common.GetInstance().GetMqqtClient().Publish("s/us", 2, false, createTemperatureMeasurement(temp, ""))
		time.Sleep(3 * time.Second)
	}
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
