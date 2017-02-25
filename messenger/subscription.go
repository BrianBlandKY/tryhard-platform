package messenger

// Subscription
type Subscription interface {
	Subject() string
	Queue() string
	IsValid() bool
	AutoUnsubscribe(max int) error
	Unsubscribe() error
}
