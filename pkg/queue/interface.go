package queue

type Producer interface {
	Publish(topic string, data []byte) error
}

type Consumer interface {
	Consume(topic string, handler interface{}) error
}
