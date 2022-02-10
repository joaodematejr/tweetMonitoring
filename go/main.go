package main

import (
	"crypto/tls"
	"encoding/json"

	"github.com/joaodematejr/tweetMonitoring/email"
	"github.com/joaodematejr/tweetMonitoring/kafka"
)

func main() {

	var emailCh = make(chan email.Email)
	var msgChan = make(chan *ckafka.Message)

	d := gomail.NewDialer("smtp.gmail.com", 587, "", "")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	es := email.NewMainSender()
	es.From = "email@email.com.br"
	es.Dailer = d

	go es.Send(msgChan)

	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": "kafka:9094",
		"client.id":         "emailapp",
		"group.id":          "emailapp",
	}
	topics := []string{"email"}
	consumer := kafka.NewConsumer(configMap, topics)
	go consumer.Consume(msgChan)

	for msg := range msgChan {
		var input email.Email
		json.Unmarshal(msg.Value, &input)
		emailCh <- input
	}
}
