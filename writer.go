package main

// NSQ Producer

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/bitly/go-nsq"
)

var counter int

func WriteMessage(numberOfMessages int, intervalInSeconds int) {
	localConfiguration := GetConfig()
	cfg := nsq.NewConfig()
	cfg.MaxInFlight = localConfiguration.MaxInFlight

	w, _ := nsq.NewProducer(localConfiguration.NSQDFullPath, cfg)
	_publishMessage := func() {
		err := w.Publish(localConfiguration.Topic, []byte("test "+strconv.Itoa(counter)))
		if err != nil {
			log.Panic("Could not connect")
		}
		log.Printf("Publishing message %d", counter)
		counter++
	}
	stop := schedule(_publishMessage, time.Duration(intervalInSeconds)*time.Second)
	time.Sleep(time.Duration(intervalInSeconds*numberOfMessages) * time.Second)
	stop <- true
	fmt.Println("Done Writing messages")
	w.Stop()
}

// schedule function will run callback function
// until the delay time
func schedule(cb func(), delay time.Duration) chan bool {
	stop := make(chan bool)

	go func() {
		for {
			cb()
			select {
			case <-time.After(delay):
			case <-stop:
				return
			}
		}
	}()

	return stop
}
