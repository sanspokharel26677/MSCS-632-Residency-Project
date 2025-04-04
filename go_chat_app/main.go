package main

import (
	"fmt"
	"time"
)

/////////////////////////////////////////////////////////////
// This is the main entry point of the Go Chat Application //
// It sets up the ChatServer and starts two interactive    //
// user sessions concurrently using goroutines.            //
/////////////////////////////////////////////////////////////

func main() {
	// Create the main channel for chat communication
	messageChannel := make(chan Message)

	// Initialize the chat server
	server := ChatServer{
		Messages: []Message{},
		Input:    messageChannel,
	}

	// Start the chat server to listen for messages
	go server.Start()

	// Create two users and pass the channel to them
	user1 := User{ID: "User1", Output: messageChannel}
	user2 := User{ID: "User2", Output: messageChannel}

	// Start each user session in a separate goroutine
	go user1.Start()
	go user2.Start()

	// Allow chat to run for a fixed time before exiting
	// You can change this or wait for manual termination later
	time.Sleep(30 * time.Second)

	// Optional: filter history
	fmt.Println("\n=== Chat History Filters ===")
	server.FilterByUser("User1")
	server.FilterByKeyword("hello")

	// End program
	fmt.Println("\n[Server] Chat session ended.")
}
