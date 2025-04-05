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

// Start begins the message listening loop.
// It continuously listens for new messages and displays them.
func (c *ChatServer) StartServer(users map[string]*User) {
	// Start a goroutine for each user to listen on their channel
	for _, user := range users {
		go func(u *User) {
			for {
				select {
				case msg, ok := <-u.Input:
					if !ok {
						// Channel closed, exit goroutine
						return
					}
					// Process the message received on the user's channel

					time.Sleep(100 * time.Millisecond)
					fmt.Printf("[%s] Received message: %v\n", u.ID, msg)
				}
			}
		}(user)
	}
}

// FilterByUser prints all messages from a specific user.
func (c *ChatServer) FilterByUser(userID string) {
	fmt.Printf("\n--- Messages from %s ---\n", userID)
	for _, msg := range c.Messages {
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

// currentTimestamp returns the current time.
// It is used when creating a new Message.
func currentTimestamp() time.Time {
	return time.Now()
}
