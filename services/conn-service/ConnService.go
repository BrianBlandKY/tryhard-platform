package main

import (
	"log"
	"time"
	"tryhard-platform/config"
	mess "tryhard-platform/messenger"
	"tryhard-platform/model"
)

/*
	This service is responsibile for 3 features:
	- Tracking Connected Services (self included)
	- Managing heartbeats for services and clients
*/
type connService struct {
	mess.Node
	// track connected users?
	services    map[string]model.Service
	connCh      chan mess.Command
	discCh      chan mess.Command
	heartbeatCh chan mess.Command
}

func (s *connService) Run() {
	s.SetHandler(s.handler)
	s.SetProcessor(s.processor)
	s.Node.Run()
}

func (s *connService) handler(n *mess.Node) {
	log.Println("Service Handler")
	for {
		select {
		case m := <-n.SubscriptionChan:
			switch m.Action {
			case mess.CONNECT:
				s.connCh <- m
			case mess.DISCONNECT:
				s.discCh <- m
			case mess.HEARTBEAT:
				s.heartbeatCh <- m
			default:
				log.Println("Unsupported action", m.Action)
			}
		}
	}
}

func (s *connService) processor() {
	log.Println("Service Processor")
	for {
		select {
		case cmd := <-s.connCh:
			log.Println(s.Config.Service.Name, "New connection established", cmd)
			// Add service to collection
			var service model.Service
			_ = cmd.Deserialize(&service)

			service.Status = int64(mess.CONNECTED)
			s.services[service.ID] = service

			_ = cmd.Serialize(service)

			// Reply
			s.Messenger.Reply(cmd)

		case cmd := <-s.discCh:
			// Remove service from collection
			var service model.Service
			_ = cmd.Deserialize(&service)

			service.Status = int64(mess.DISCONNECTED)
			s.services[service.ID] = service

			_ = cmd.Serialize(service)

			// Reply
			s.Messenger.Reply(cmd)

		case cmd := <-s.heartbeatCh:
			// Heartbeat Reply
			var heartbeat mess.Heartbeat
			_ = cmd.Deserialize(&heartbeat)

			heartbeat.ServiceReceivedTime = time.Now()
			diff := time.Since(heartbeat.NodeStartTime)
			heartbeat.UploadLatency = diff.Seconds() * 1000

			_ = cmd.Serialize(heartbeat)
			_ = s.Messenger.Reply(cmd)
		}
	}
}

func newConnectionService(cfg config.Config) (s connService) {
	s = connService{
		Node:        mess.DefaultNode(cfg),
		heartbeatCh: make(chan mess.Command),
		connCh:      make(chan mess.Command),
		discCh:      make(chan mess.Command),
		services:    make(map[string]model.Service),
	}
	return s
}
