package main

import (
	"fmt"
)

// User represents a chat participant with a unique ID and
// a channel to send messages to the main chat server.
type User struct {
	ID    string       // Unique identifier for the user
	Input chan Message // Channel used to receive messages
}

func AddUser(users map[string]*User, id string, input chan Message) (*User, error) {
	if _, exists := users[id]; exists {
		return nil, fmt.Errorf("user with ID %s already exists", id)
	}

	user := &User{
		ID:    id,
		Input: input, // Input channel will be set when used in the chat server
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
