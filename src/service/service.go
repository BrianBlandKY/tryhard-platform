package service

import (
	"fmt"
	cfg "tryhard-platform/src/config"
	log "tryhard-platform/src/logging"
	msg "tryhard-platform/src/messenger"
)

// Service - Base Implementation for Services
type Service interface {
	log.Logger
	Config() cfg.Config
	Status() msg.Status
	Start() <-chan msg.Command
	Stop()
	Request(cmd msg.Command, response *msg.Command) error
	Reply(cmd msg.Command) error
	Execute()
}

type service struct {
	log.Logger
	config      cfg.Config
	messenger   msg.Messenger
	subStopChan chan bool
	subChan     chan msg.Command
	sub         msg.Subscription
}

func (svc *service) Config() cfg.Config {
	return svc.config
}

func (svc *service) Status() msg.Status {
	return svc.messenger.Status()
}

func (svc *service) Start() (ch <-chan msg.Command) {
	if svc.Status() == msg.CONNECTED {
		svc.Println("Service connected.")
		return
	}

	// connect
	svc.Println(svc.config.Service.Name, "Connecting to", svc.config.Platform.Address)
	err := svc.messenger.Connect(svc.config.Platform.Address)
	if err != nil {
		// panic here but once connection is established
		// we'll utilize the nats auto-reconnect to stay running
		panic(err)
	}
	svc.Println(svc.config.Service.Name, "Connected to", svc.config.Platform.Address)

	// subscribe
	svc.subChan = make(chan msg.Command)
	key := svc.messenger.Key(svc.config.Service.Name, "*")
	sub, err := svc.messenger.SubscribeChan(key, svc.subChan)
	if err != nil {
		panic(fmt.Errorf("error creating channel %v", err))
	}

	svc.sub = sub
	svc.Println(svc.config.Service.Name, "Subscribed to", key)

	// start relay
	// return relay chan
	relayCh := make(chan msg.Command)
	go svc.subscriptionWrapperRelay(relayCh)

	return relayCh
}

func (svc *service) Stop() {
	if svc.Status() == msg.CLOSED {
		svc.Println("Service is closed.")
		return
	}

	if svc.Status() == msg.DISCONNECTED {
		svc.Println("Service is disconnected.")
		return
	}

	// stop relay
	svc.subStopChan <- true

	key := svc.messenger.Key(svc.config.Service.Name, "*")
	svc.sub.Unsubscribe()
	svc.Println(svc.config.Service.Name, "Unsubscribed from", key)

	svc.messenger.Close()
	svc.Println(svc.config.Service.Name, "Disconnected from", svc.config.Platform.Address)
}

// TODO: automatic retry?
func (svc *service) Request(cmd msg.Command, response *msg.Command) error {
	return svc.messenger.Request(cmd, response)
}

// TODO: automatic retry?
func (svc *service) Reply(cmd msg.Command) error {
	return svc.messenger.Reply(cmd)
}

func (svc *service) Execute() {
	panic("this should be overwritten from core service.")
}

func (svc *service) subscriptionWrapperRelay(ch chan<- msg.Command) {
	for {
		select {
		case m := <-svc.subChan:
			// command audit
			// track commands received
			// response time?
			// relay
			ch <- m
		case _ = <-svc.subStopChan:
			break
		}
	}
}

func DefaultService(cfg cfg.Config) Service {
	svc := &service{
		config:      cfg,
		Logger:      log.DefaultLogger(),
		messenger:   msg.DefaultMessenger(),
		subStopChan: make(chan bool),
	}
	return svc
}
