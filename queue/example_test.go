package queue

import (
	"encoding/base64"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/yqchilde/pkgs/queue/nats"
)

func TestNats(t *testing.T) {
	var (
		addr  = "nats://localhost:4222"
		topic = "hello"
	)
	producer := nats.NewProducer(addr)
	consumer := nats.NewConsumer(addr)

	published := make(chan struct{})
	received := make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		for {
			select {
			case <-published:
				time.Sleep(3 * time.Second)
				if err := producer.Publish(topic, []byte("hello nats")); err != nil {
					t.Error(err)
				}
				t.Log("producer handler publish msg: ", "hello nats")

			case <-received:
				wg.Done()
				break
			}
		}
	}()
	go func() {
		for {
			select {
			default:
				handler := func(message []byte) error {
					decodeMessage, _ := base64.StdEncoding.DecodeString(strings.Trim(string(message), "\""))
					t.Log("consumer handler receive msg: ", string(decodeMessage))
					received <- struct{}{}
					wg.Done()
					return nil
				}
				if err := consumer.Consume(topic, handler); err != nil {
					t.Error(err)
				}
				time.Sleep(5 * time.Second)
			}
		}
	}()

	published <- struct{}{}
	wg.Wait()
}
