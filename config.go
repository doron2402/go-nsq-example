package main

type configuration struct {
	Channel      string
	Topic        string
	NSQHost      string
	NSQPort      string
	NSQDFullPath string
	MaxInFlight  int
}

// Get configuration
func GetConfig() configuration {
	Config := configuration{
		Channel:     "test_channel",
		Topic:       "test_topic",
		NSQHost:     "127.0.0.1",
		NSQPort:     "4150",
		MaxInFlight: 10,
	}
	Config.NSQDFullPath = Config.NSQHost + ":" + Config.NSQPort
	return Config
}
