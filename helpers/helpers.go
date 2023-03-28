package helpers

import (
	"batch-acctstatement/config"
	"encoding/json"
	"log"
	"os"
)

type (
	Message struct {
		Container string      `json:"containerID"`
		Type      string      `json:"type"`
		Payload   interface{} `json:"payload"`
	}
)

func SendSignKill() {
	hostname, _ := os.Hostname()
	log.Print("Exiting scripts...")
	PubLogMsg("INFO", "Exiting scripts...")

	c, err := config.ConnectNats()
	if err != nil {
		log.Print(err)
	}
	c.Stan.Publish("STOP_CONTAINER", []byte(hostname))

	c.Stan.Close()
	c.Nats.Close()
}

func PubLogMsg(logType string, msg string) {
	hostname, _ := os.Hostname()
	logMsg := &Message{
		Container: hostname,
		Type:      logType,
		Payload:   msg,
	}
	payload, _ := json.Marshal(logMsg)
	c, err := config.ConnectNats()
	if err != nil {
		log.Print(err)
	}
	c.Stan.Publish("LOGGING", payload)
}
