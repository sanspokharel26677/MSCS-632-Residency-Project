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
	UserID  string    // ID of the user who sent the message
	Content string    // Actual message text
	Time    time.Time // Timestamp of when message was sent
}

// ChatServer handles storing and displaying messages.
type ChatServer struct {
	Messages []Message        // Slice to store chat history
	Input    chan Message     // Channel to receive messages
}

// Start begins the message listening loop.
// It continuously listens for new messages and displays them.
func (c *ChatServer) Start() {
	for msg := range c.Input {
		c.Messages = append(c.Messages, msg) // Store message
		fmt.Printf("[%s][%s]: %s\n", msg.Time.Format("15:04:05"), msg.UserID, msg.Content)
	}
}

// FilterByUser prints all messages from a specific user.
func (c *ChatServer) FilterByUser(userID string) {
	fmt.Printf("\n--- Messages from %s ---\n", userID)
	for _, msg := range c.Messages {
		if msg.UserID == userID {
			fmt.Printf("[%s] %s\n", msg.Time.Format("15:04:05"), msg.Content)
		}
	}
}

// FilterByKeyword prints all messages containing the given keyword.
func (c *ChatServer) FilterByKeyword(keyword string) {
	fmt.Printf("\n--- Messages containing '%s' ---\n", keyword)
	for _, msg := range c.Messages {
		if strings.Contains(strings.ToLower(msg.Content), strings.ToLower(keyword)) {
			fmt.Printf("[%s][%s]: %s\n", msg.Time.Format("15:04:05"), msg.UserID, msg.Content)
		}
	}
}

// currentTimestamp returns the current time.
// It is used when creating a new Message.
func currentTimestamp() time.Time {
	return time.Now()
}
