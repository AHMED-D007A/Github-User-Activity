package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Usage: ./github_activity <UserName>")
		return
	}

	userName := os.Args[1]
	url := fmt.Sprintf("https://api.github.com/users/%s/events/public", string(userName))

	response, err := http.Get(url)
	if err != nil {
		log.Fatalln(err.Error())
	}

	if response.Status == "404 Not Found" {
		log.Fatalln("Invalid UserName Input.")
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err.Error())
	}

	var events []Event

	err = json.Unmarshal(responseData, &events)
	if err != nil {
		log.Fatalln(err.Error())
	}

	for _, event := range events {
		switch event.Type {
		case "CreateEvent":
			createEvent(event)
		case "PushEvent":
			pushEvent(event)
		case "WatchEvent":
			watchEvent(event)
		}
	}
}
