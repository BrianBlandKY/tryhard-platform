package messenger

import (
	"bufio"
	"fmt"
	"os"
	"time"
	cfg "tryhard-platform/config"
	log "tryhard-platform/logging"
	"tryhard-platform/model"
)

/*
Node - Base implementation for services.
Features:
	- Connect
	- Disconnect
	- Heartbeating
*/
type Node struct {
	log.Logger
	Status             Status     `json:"status"`
	Config             cfg.Config `json:"config"`
	Messenger          Messenger
	Subscription       Subscription
	SubscriptionChan   chan Command
	heartbeatStopCh    chan bool
	heartbeatTicker    *time.Ticker
	heartbeatTimeCh    <-chan time.Time
	statusCh           chan Status
	serviceHandlerFn   func(*Node)
	serviceProcessorFn func()
	consoleFunc        func()
}

func (n *Node) Close() {
	n.Println("Closing...")

	// send DISCONNECT Command to ConnectionService
	cmd := Command{
		Service: Services.Connection,
		Action:  DISCONNECT,
	}

	cmd.Serialize(model.Service{
		ID:   n.Config.Service.ID,
		Name: n.Config.Service.Name,
	})

	var resCommand Command
	err := n.Messenger.Request(cmd, &resCommand)
	if err != nil {
		n.Printf("error %v", err)
	}
	n.Messenger.Close()
	n.heartbeatStopCh <- true
}

// SetHandler Define Handler?
func (n *Node) SetHandler(serviceHandlerFn func(n *Node)) {
	n.serviceHandlerFn = serviceHandlerFn
}

// SetProcessor Define Procesor?
/*
	Why do we need to set a process? This allows us to start the processor
	after connection but before the handler. This in turn, allows the service
	to listen to itself for commands. See ConnectionService, where we needed to
	start the processor first to listen for Connection command.
*/
func (n *Node) SetProcessor(serviceProcessorFn func()) {
	n.serviceProcessorFn = serviceProcessorFn
}

// SetConsole override the default console input scanner.
// Useful for test clients
func (n *Node) SetConsole(consoleFn func()) {
	n.consoleFunc = consoleFn
}

// Run Function
func (n *Node) Run() {
	// connect
	n.connect()

	// subscribe
	n.subscribe(n.Config.Service.Name)

	// processor
	if n.serviceProcessorFn != nil {
		go n.serviceProcessorFn()
		n.Println(n.Config.Service.Name, "Service Processor Running...")
	}

	// heartbeat processor
	go n.internalProcessor()

	// send connection command to CONNECTION service
	// TODO: Find better way to trigger connection command. timer not recommended
	time.AfterFunc(1*time.Second, func() { n.internalConnectionCommand() })

	go n.consoleFunc()

	// service listener
	if n.serviceHandlerFn != nil {
		n.Println(n.Config.Service.Name, "Service Listener Running...")
		n.serviceHandlerFn(n)
	}
	n.Println("Service done. Listener broken.")
}

func (n *Node) connect() {
	n.Println(n.Config.Service.Name, "Connecting to", n.Config.Platform.Address)
	err := n.Messenger.Connect(n.Config.Platform.Address)
	if err != nil {
		panic(err)
	}
	n.Println(n.Config.Service.Name, "Connected to", n.Config.Platform.Address)
}

// Subscribe Each Node/Service will have a single subscription tied to the [Name].*
func (n *Node) subscribe(serviceName string) {
	key := n.Messenger.Key(serviceName, "*")
	recvSub, err := n.Messenger.SubscribeChan(key, n.SubscriptionChan)
	if err != nil {
		n.Printf("error creating channel %v \r\n", err)
	}

	n.Subscription = recvSub
}

func (n *Node) internalProcessor() {
	// service loop for
	for {
		select {
		case <-n.heartbeatTimeCh:
			heartbeat := Heartbeat{
				NodeStartTime: time.Now(),
			}

			cmd := Command{
				Service: Services.Connection,
				Action:  HEARTBEAT,
			}

			_ = cmd.Serialize(heartbeat)

			n.Verboseln(n.Config.Settings.VerboseHeartbeat, n.Config.Service.Name, "HEARTBEAT: Sending.")

			var resCmd Command
			err := n.Messenger.Request(cmd, &resCmd)
			if err != nil {
				n.statusCh <- n.Messenger.Status()
				n.Verboseln(n.Config.Settings.VerboseHeartbeat, n.Config.Service.Name, "HEARTBEAT: Failed.", err)
			}

			resCmd.Deserialize(&heartbeat)

			heartbeat.NodeEndTime = time.Now()
			diff := time.Since(heartbeat.ServiceReceivedTime)
			heartbeat.DownloadLatency = diff.Seconds() * 1000

			n.Verbosef(n.Config.Settings.VerboseHeartbeat, "%s HEARTBEAT: Received upload:%fms download:%fms", n.Config.Service.Name, heartbeat.UploadLatency, heartbeat.DownloadLatency)
			go n.updateStatus()

		case status := <-n.statusCh:
			n.Status = status
			if n.Status == CONNECTED && n.heartbeatTicker == nil {
				// heartbeat ticker
				ticker := time.NewTicker(2 * time.Second)
				n.heartbeatTicker = ticker
				n.heartbeatTimeCh = ticker.C
			}

		// update messenger status
		case <-n.heartbeatStopCh:
			n.heartbeatTicker.Stop()
			return
		default:
			//log.Println("do something inside node")
		}
	}
}

func (n *Node) internalConnectionCommand() {
	// send CONNECT Command to ConnectionService
	cmd := Command{
		Service: Services.Connection,
		Action:  CONNECT,
	}

	cmd.Serialize(model.Service{
		ID:   n.Config.Service.ID,
		Name: n.Config.Service.Name,
	})

	var resCommand Command
	err := n.Messenger.Request(cmd, &resCommand)
	if err != nil {
		// Auto-Reconnect?
		panic(fmt.Sprintf("Failed to connect to a Connection Service. Service Abandoned. %v \r\n", err))
	}
	n.updateStatus()
}

/*
	Future Features:
		- Print Connection Status
			- Messenger
			- Node (should always match Messenger?)
		- Print Heartbeat Stats
			- Current Heartbeat Latency
			- Current Average Heartbeat Latency
			- Total Heartbeats
		- Print Config
			- Id
			- Name
			- Address
		- Enable/Disable/Config Heartbeat Output
		- Enable/Disable/Config Message Output
*/
func (n *Node) internalConsoleScanner() {
	// Show some cool output?
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		command := scanner.Bytes()
		n.Println("command:", command)

		switch cmd := string(command[:len(command)]); cmd {
		// case "connect":
		// 	go c.connect()
		default:
			n.Println("unrecognized command", cmd)
		}
	}
}

func (n *Node) updateStatus() {
	n.statusCh <- n.Messenger.Status()
}

// DefaultNode Gets the default node implementation using the default messenger
func DefaultNode(cfg cfg.Config) Node {
	node := Node{
		Config:           cfg,
		Logger:           log.DefaultLogger(),
		Messenger:        defaultMessenger(),
		SubscriptionChan: make(chan Command),
		heartbeatStopCh:  make(chan bool),
		heartbeatTimeCh:  make(chan time.Time),
		statusCh:         make(chan Status),
	}
	node.SetConsole(node.internalConsoleScanner)
	return node
}
