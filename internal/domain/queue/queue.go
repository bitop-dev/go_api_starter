package queue

// Publisher defines minimal behavior for an async message queue provider.
type Publisher interface {
	Publish(subject string, payload []byte) error
}
