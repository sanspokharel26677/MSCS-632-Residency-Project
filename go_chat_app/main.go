package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

/////////////////////////////////////////////////////////////
// This is the main entry point of the Go Chat Application //
// It sets up the ChatServer and starts two interactive    //
// user sessions concurrently using goroutines.            //
/////////////////////////////////////////////////////////////

func main() {
	// Initialize the chat server
	server := ChatServer{
		Users:    map[string]*User{}, // Initialize the users map
		Messages: []Message{},
	}

	// Start the chat server to listen for messages
	var users map[string]*User
	users = make(map[string]*User)

	AddUser(users, "User1", make(chan Message))
	AddUser(users, "User2", make(chan Message))
	server.StartServer(users) // Start the server to listen for messages
	for {
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Would you like to create a new user? (yes/no): ")
		response, _ := reader.ReadString('\n')
		if strings.ToLower(strings.TrimSpace(response)) != "yes" && strings.ToLower(strings.TrimSpace(response)) != "no" {
			fmt.Println("Invalid response, please enter 'yes' or 'no'.")
		}
		if strings.ToLower(strings.TrimSpace(response)) == "yes" {
			fmt.Print("Enter a username for the new user: ")
			newUserName, _ := reader.ReadString('\n')
			newUserName = strings.TrimSpace(newUserName)

			// Check if the user already exists
			if _, exists := users[newUserName]; exists {
				fmt.Printf("User %s already exists, please choose a different username.\n", newUserName)
				continue
			}

			// Create a new user and add to the server
			newUser, err := AddUser(users, newUserName, make(chan Message))
			if err != nil {
				fmt.Println("Error creating user:", err)
				continue
			}
			fmt.Printf("New user %s created successfully.\n", newUser.ID)
		}

		fmt.Print("Enter sender username: ")
		userFrom, _ := reader.ReadString('\n')
		userFrom = strings.TrimSpace(userFrom)

		fmt.Print("Enter recipient username: ")
		userTo, _ := reader.ReadString('\n')
		userTo = strings.TrimSpace(userTo)

		fmt.Print("Enter message (or 'exit' to quit): ")
		message, _ := reader.ReadString('\n')
		message = strings.TrimSpace(message)

		// Check if both users exist
		_, fromExists := users[userFrom]
		_, toExists := users[userTo]

		if !fromExists {
			fmt.Printf("Error: User %s does not exist\n", userFrom)
			continue
		}
		if !toExists {
			fmt.Printf("Error: User %s does not exist\n", userTo)
			continue
		}
		var newMessage Message
		newMessage.UserIDFrom = userFrom
		newMessage.UserIDTo = userTo
		newMessage.Content = message
		userToSendTo := users[userTo]    // Get the recipient user object
		userToSendTo.Input <- newMessage // Send the message to the recipient's input channel

		newMessage.Time = time.Now() // Set the current time for the message
		time.Sleep(500 * time.Millisecond)
	}

}
