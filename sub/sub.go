package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"crypto/tls"
	"crypto/x509"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"io/ioutil"
)

var f MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Println("----------")
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

var onConnect MQTT.OnConnectHandler = func(client MQTT.Client) {
	fmt.Println("--CONNECTED--")
}

var onConnectionLost MQTT.ConnectionLostHandler = func(client MQTT.Client, errmsg error) {
	fmt.Println("--DISCONNECTED--")
	fmt.Printf("error: %s\n", errmsg.Error())
}

var (
	serverAddrOpt = flag.String("s", "127.0.0.1", "MQTT server address")
	clientIDOpt   = flag.String("i", "pub", "MQTT client ID")
	topicOpt      = flag.String("t", "topic", "MQTT topic")
	userNameOpt   = flag.String("u", "user", "MQTT user")
	userPassOpt   = flag.String("p", "user", "MQTT pass")
)

func main() {
	flag.Parse()

	fmt.Printf("Server: %s\n", *serverAddrOpt)
	fmt.Printf("Client ID: %s\n", *clientIDOpt)
	fmt.Printf("Topic: %s\n", *topicOpt)
	fmt.Printf("Username: %s\n", *userNameOpt)
	fmt.Printf("Password: %s\n", *userPassOpt)

	//broker := "tcp://" + *serverAddrOpt + ":1883"
	broker := "ssl://" + *serverAddrOpt + ":8883"
	opts := MQTT.NewClientOptions().AddBroker(broker)
	opts.SetClientID(*clientIDOpt)
	//opts.SetDefaultPublishHandler(f)
	opts.SetOnConnectHandler(onConnect)
	opts.SetConnectionLostHandler(onConnectionLost)
	opts.CleanSession = true

	opts.SetUsername(*userNameOpt)
	opts.SetPassword(*userPassOpt)

	certpool := x509.NewCertPool()
	pemCerts, err := ioutil.ReadFile("certs/ca/ca.crt")

	if err != nil {
		panic(err)
	}

	certpool.AppendCertsFromPEM(pemCerts)

	cer, err := tls.LoadX509KeyPair("certs/server.crt", "certs/server.key")
	if err != nil {
		panic(err)
	}

	tlsConfig := &tls.Config{
		RootCAs:      certpool,
		Certificates: []tls.Certificate{cer},
	}

	opts.SetTLSConfig(tlsConfig)

	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	subHandler := func(c MQTT.Client, msg MQTT.Message) {
		fmt.Printf("Get new item: \nTopic: %v, Payload:%v\n", msg.Topic(), msg.Payload())
	}

	if token := c.Subscribe(*topicOpt, 2, subHandler); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	time.Sleep(3600 * time.Second)

	if token := c.Unsubscribe(*topicOpt); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}
}
