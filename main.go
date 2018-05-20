package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func parseArgs(args []string) map[string]string {
	temp := make(map[string]string)
	log.Printf("Length: %d", len(args))
	for i := 0; i <= (len(args) - 1); i++ {
		temp[args[i]] = args[(i + 1)]
	}
	return temp
}

func main() {
	fmt.Printf("Starting...")
	var numberOfMessages int = 3
	var intervalInSeconds int = 1

	args := parseArgs(os.Args[1:])

	if args["action"] == "publish" {
		// Write message

		// check for number of messages
		if _, ok := args["messages"]; ok {
			value, err := strconv.Atoi(args["messages"])
			if err != nil {
				log.Panic("Error parsing number of messages")
			}
			numberOfMessages = value
		}
		// check for interval in seconds
		if _, ok := args["interval"]; ok {
			value, err := strconv.Atoi(args["interval"])
			if err != nil {
				log.Panic("Error parsing interval in seconds")
			}
			intervalInSeconds = value
		}
		fmt.Printf("Publish messages")
		WriteMessage(numberOfMessages, intervalInSeconds)
	} else if args["action"] == "consume" {
		fmt.Printf("Read messages")
		ReadMessage()
	} else {
		log.Panic("Unknown `action` (valid are `publish` or `consume`")
	}

}
