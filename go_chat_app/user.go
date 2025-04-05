package main

import (
	"fmt"
)

// User represents a chat participant with a unique ID and
// a channel to send messages to the main chat server.
type User struct {
	ID               string       // Unique identifier for the user
	Input            chan Message // Channel used to receive messages
	SentMessages     []Message    // Store messages sent by this user, if needed
	ReceivedMessages []Message    // Store messages received by this user, if needed
}

func AddUser(users map[string]*User, id string, input chan Message) (*User, error) {
	if _, exists := users[id]; exists {
		return nil, fmt.Errorf("User with ID %s already exists", id)
	}

	user := &User{
		ID:               id,
		Input:            input,       // Input channel will be set when used in the chat server
		SentMessages:     []Message{}, // Initialize sent messages slice
		ReceivedMessages: []Message{}, // Initialize received messages slice
	}
	users[id] = user
	return user, nil
}

func RemoveUser(users map[string]*User, id string) {
	delete(users, id)
}

func GetUser(users map[string]*User, id string) (*User, bool) {
	user, exists := users[id]
	return user, exists
}

func (u *User) StartRouting() {
	// Start a goroutine to listen for messages on the user's input channel
	go func() {
		for {
			select {
			case msg, ok := <-u.Input:
				if !ok {
					// Channel closed, exit goroutine
					return
				}
				// Process the message received on the user's channel
				fmt.Printf("%s | %s -> %s: %s\n", msg.Time.Format("15:04:05"), msg.UserIDFrom, msg.UserIDTo, msg.Content)
			}
		}
	}()
}
