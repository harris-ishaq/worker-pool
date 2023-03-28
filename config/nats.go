package config

import (
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
)

type (
	DNats struct {
		Nats *nats.Conn
		Stan stan.Conn
	}
)

func ConnectNats() (*DNats, error) {
	var now = time.Now()
	natsConfig := GetNatsConfig()
	nc, err := natsConfig.Connect()
	if err != nil {
		log.Printf("Error connecting to nats server, cause: %+v \n", err)
		return nil, err
	}
	sc, err := stan.Connect(CLUSTERID, fmt.Sprint(HOSTNAME, "-", now.UnixNano()), stan.NatsConn(nc))
	if err != nil {
		log.Printf("Error connecting to stan service, cause: %+v \n", err)
		nc.Close()
		return nil, err
	}
	var dnats = &DNats{
		Nats: nc,
		Stan: sc,
	}
	return dnats, nil
}

func (d *DNats) StanPublish(subject string, data []byte) error {
	return d.Stan.Publish(subject, data)
}

func (d *DNats) Disconnect() {
	d.Stan.Close()
	d.Nats.Close()
}

func GetNatsConfig() nats.Options {
	opts := nats.GetDefaultOptions()
	opts.Url = NATSURL
	opts.MaxPingsOut = 10000
	opts.MaxReconnect = 10000
	opts.ReconnectWait = 30 * time.Second
	opts.PingInterval = 5 * time.Second
	opts.Verbose = true
	opts.AllowReconnect = true
	return opts
}
