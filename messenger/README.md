## Dialect - Abstract Interface for Command Communication

### Supported Protocols
- NATS 

### Interface

#### Dialect
- Connect(url string) (Connection, err)
- ConnectEncoding(url string, encoding string) (Connection, err)

#### Connection
- PublishCommand(command Command)
- Publish(key string, command Command)
- Subscribe(key string, func(command Command))
- SubscribeChan(key string, *command chan)
- Request(key string, command Command, response Command)
- RequestWithTimeout(key string, command Command, response Command, seconds int)
- Close()

#### Subscription
- Unsubscribe()