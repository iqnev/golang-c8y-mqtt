# golang-c8y-mqtt

A simple example of an MQTT client for the C8Y. It creates a new device, subscrubes for certain topic, creates  an MQTT Smart templates and pushes measurements periodically into the platform

## Requirements:

`go v1.15`

## Quick Start

1. Clone the repository
```shell
git clone https://github.com/iqnev/golang-c8y-mqtt.git
cd golang-c8y-mqtt
```
2. Configure the Cumulocity  settings
You have to create a json file with the following template
```json
{
    "C8Y_SEREVR_URL" : "tcp://cumulocity.com:1883",
    "DEVICE_NAME" : "deviceName",
    "C8Y_TENENT" : "myTenent",
    "C8Y_USERNAME" : "myuser@mail.com",
    "C8Y_PASSWORD" : "myPassword",
    "CLIENT_ID"     : "clientID",
    "DEVICE_OPERATIONS" : "c8y_Restart, c8y_SoftwareList.... [myOperations]",
    "REVESION" : "deviceRevesion",
    "SERIAL_NUMBER" : "deviceNumber",
    "HARDWARE_MODEL" : "deviceModel",
    "REQUIRED_INTERVAL" : "deviceRequiredInterval",
    "SMART_REST_TEMPLATE" : "deviceCustomMqttTemplate",
    "SMART_REST_TEMPLATE_ID": "templateID"	
}
```

The file name of the above configuration is important because you have the ability to create multiple configurations depending on your environment. The name convention is: `[environment]_config.json`  By default you need to create a dev configuration: `dev_config.json`

3. Run the device application

`go run main.go`

and you can pass the given environment as command line Arguments

`go run main.go [environment]` 
