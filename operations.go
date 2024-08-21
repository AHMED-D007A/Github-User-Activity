package main

import "fmt"

func createEvent(event Event) {
	fmt.Printf("- Created a %v %v at %v\n", event.Payload.Ref_Type, event.Payload.Ref, event.Repo.Name)
}

func pushEvent(event Event) {
	fmt.Printf("- Pushed %v commit(s) to %v\n", len(event.Payload.Commit), event.Repo.Name)
}

func watchEvent(event Event) {
	if event.Payload.Action == "started" {
		fmt.Printf("- Watched and Started the repository at %v\n", event.Repo.Name)
	} else {
		fmt.Printf("- Watched the repository at %v\n", event.Repo.Name)
	}
}
