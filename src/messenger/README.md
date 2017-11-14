## Dialect - Abstract Interface for Command Communication

### Supported Protocols
- NATS 

### Interface

#### Messenger
- Key(cmd Command)
- Connect(url, id string) error
- Publish(command Command) error
- Subscribe(key string, func(command Command)) (Subscription, error)
- Request(command Command, response *Command)
- RequestTimeout(command Command, response *Command, seconds int)
- Reply(command Command) error
- BindRecvChan(key string, ch chan *Command) (Subscription, error)
- BindSendChan(key string, ch chan *Command) error
- Close()
- ID() string
- Status() Status
- IsConnected() bool
- IsClosed() bool

#### Subscription
- Subject() string
- Queue() string
- IsValid() bool
- AutoUnsubscribe(max int) error
- Unsubscribe() error