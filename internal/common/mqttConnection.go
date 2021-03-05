package common

import (
	"sync"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"
)

type Instance struct {
	mqttCl mqtt.Client
}

var (
	Singletone    *Instance
	once          sync.Once
	clientOptions *mqtt.ClientOptions
)

func (client *Instance) GetMqqtClient() mqtt.Client {
	return client.mqttCl
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	log.Warn("Device is Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	log.Warn("Connect lost: %v", err)
}

func GetInstance() *Instance {
	once.Do(func() {
		client := mqtt.NewClient(clientOptions)
		Singletone = &Instance{mqttCl: client}

		if token := client.Connect(); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}
	})

	return Singletone
}

func CloseConnection() {
	log.Info("Disconnecting Connection....")
	GetInstance().mqttCl.Disconnect(255)

}

func InitClientOptions(conf Configuration) {
	opts := mqtt.NewClientOptions()

	opts.AddBroker(conf.C8Y_SEREVR_URL)
	opts.SetUsername(conf.C8Y_TENENT + "/" + conf.C8Y_USERNAME)
	opts.SetPassword(conf.C8Y_PASSWORD)
	opts.SetClientID(conf.CLIENT_ID)
	//opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler

	clientOptions = opts
}
