package main

import (
	"flag"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"time"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
)

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

	broker := "ssl://" + *serverAddrOpt + ":8883"
	opts := MQTT.NewClientOptions().AddBroker(broker)
	opts.SetClientID(*clientIDOpt)
	opts.SetUsername(*userNameOpt)
	opts.SetPassword(*userPassOpt)
	opts.SetOnConnectHandler(onConnect)
	opts.SetConnectionLostHandler(onConnectionLost)

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
		RootCAs: certpool,
		Certificates: []tls.Certificate{cer},
	}

	opts.SetTLSConfig(tlsConfig)

	c := MQTT.NewClient(opts)
	token := c.Connect()
	if token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	messageFmt := "msg #%d from " + *clientIDOpt + " (%s)"
	for i := 0; i < 3600; i++ {
		t := time.Now()
		text := fmt.Sprintf(messageFmt, i, t)
		fmt.Println(text)
		token := c.Publish(*topicOpt, 2, false, text)
		token.Wait()
		time.Sleep(10000 * time.Millisecond)
	}
	c.Disconnect(250)
}


