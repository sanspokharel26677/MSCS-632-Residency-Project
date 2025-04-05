package main

import (
	"fmt"
	"strings"
	"time"
)

////////////////////////////////////////////////////////////
// This file defines the Message struct and ChatServer.   //
// ChatServer stores and displays messages, and supports  //
// filtering by user and keyword.                         //
////////////////////////////////////////////////////////////

// Message holds the content sent by a user along with metadata.
type Message struct {
	UserIDFrom string    // ID of the user who sent the message
	UserIDTo   string    // ID of the user to whom the message is directed
	Content    string    // Actual message text
	Time       time.Time // Timestamp of when message was sent
}

// ChatServer handles storing and displaying messages.
type ChatServer struct {
	Users    map[string]*User // Map of users to allow user management
	Messages []Message        // Slice to store chat history
}

// FilterByUser prints all messages from a specific user.
func (c *ChatServer) FilterByUser(userID string) {
	// Check if the user exists in the chat server
	if _, exists := c.Users[userID]; !exists {
		fmt.Printf("User with ID %s does not exist.\n", userID)
		return // Exit if the user does not exist
	}
	var sentMessages = c.Users[userID].SentMessages
	for _, msg := range sentMessages {
		if msg.UserIDFrom == userID || msg.UserIDTo == userID {
			fmt.Printf("%s | %s -> %s: %s\n", msg.Time.Format("15:04:05"), msg.UserIDFrom, msg.UserIDTo, msg.Content)
		}
	}

	var receivedMessages = c.Users[userID].ReceivedMessages
	for _, msg := range receivedMessages {
		if msg.UserIDFrom == userID || msg.UserIDTo == userID {
			fmt.Printf("%s | %s -> %s: %s\n", msg.Time.Format("15:04:05"), msg.UserIDFrom, msg.UserIDTo, msg.Content)
		}
	}
}

// FilterByKeyword prints all messages containing the given keyword.
func (c *ChatServer) FilterByKeyword(keyword string) {
	fmt.Printf("\n--- Messages containing '%s' ---\n", keyword)
	for _, msg := range c.Messages {
		if strings.Contains(strings.ToLower(msg.Content), strings.ToLower(keyword)) {
			fmt.Printf("%s | %s -> %s: %s\n", msg.Time.Format("15:04:05"), msg.UserIDFrom, msg.UserIDTo, msg.Content)
		}
	}
}
