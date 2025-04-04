package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//////////////////////////////////////////////////////////
// This file defines the User struct and its behavior.  //
// Each user can send messages from the CLI in real-time.//
// Messages are sent through a shared channel to the chat.//
//////////////////////////////////////////////////////////

// User represents a chat participant with a unique ID and
// a channel to send messages to the main chat server.
type User struct {
	ID     string          // Unique identifier for the user
	Output chan<- Message  // Channel used to send messages
}

// Start begins the user input loop. It reads messages from
// the command line and sends them to the chat server.
func (u *User) Start() {
	reader := bufio.NewReader(os.Stdin) // Create a reader to get input from CLI

	for {
		// Prompt the user to enter a message
		fmt.Printf("[%s] Enter message: ", u.ID)

		// Read the message input
		text, _ := reader.ReadString('\n')

		// Trim the newline character
		text = strings.TrimSpace(text)

		// If user types "/exit", break out of the loop
		if text == "/exit" {
			fmt.Printf("[%s] left the chat.\n", u.ID)
			break
		}

		// Create a Message object with the typed content
		msg := Message{
			UserID:  u.ID,
			Content: text,
			Time:    currentTimestamp(),
		}

		// Send the message through the Output channel
		u.Output <- msg
	}
}
