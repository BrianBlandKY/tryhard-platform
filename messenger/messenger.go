package messenger

import "time"

const (
	// DefaultTimeout limit set to 10 seconds
	DefaultTimeout time.Duration = time.Second * 10
)

// CommandFn Command Handler
type CommandFn func(subject, reply string, cmd Command)

// Connect with default GOBEncoder
func Connect(url string, id string) (conn Connection, err error) {
	return ConnectEncoding(url, id, GOBEncoder)
}

// ConnectEncoding connect with specific encoder
func ConnectEncoding(url string, id string, encoder string) (conn Connection, err error) {
	n := &natsDialect{}
	err = n.connect(url, id, encoder)
	if err != nil {
		return nil, err
	}

	return n, nil
}
