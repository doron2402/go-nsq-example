package main

// Consumer

import (
	"fmt"
	"log"
	"sync"

	"github.com/bitly/go-nsq"
)

func ReadMessage() {
	localConfiguration := GetConfig()
	wg := &sync.WaitGroup{}
	wg.Add(1)

	cfg := nsq.NewConfig()
	cfg.MaxInFlight = 8

	q, err := nsq.NewConsumer(localConfiguration.Topic, localConfiguration.Channel, cfg)
	if err != nil {
		log.Panic("Error creating consumer")
	}
	q.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		fmt.Printf("%s\n", message.Body)
		// wg.Done()
		return nil
	}))

	err = q.ConnectToNSQD(localConfiguration.NSQDFullPath)
	if err != nil {
		log.Panic("Could not connect to NSQD")
	}
	wg.Wait()

}
