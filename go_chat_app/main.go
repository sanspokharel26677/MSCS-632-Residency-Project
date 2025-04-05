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
	for {
		reader := bufio.NewReader(os.Stdin)
		response := AskUserForInput() // Ask the user for input

		if response == "add user" {
			fmt.Print("Enter a username for the new user: ")
			newUserName, _ := reader.ReadString('\n')
			newUserName = strings.TrimSpace(newUserName)
			_, err := CreateUser(server.Users, newUserName) // Use the helper function to create a new user

			// Create a new user and add to the server
			if err != nil {
				continue // Skip to the next iteration if there was an error adding the user
				// Skip to the next iteration if there was an error
			}
		}
		if response == "exit" {
			fmt.Println("Exiting the chat application.")
			break // Exit the loop if the user types 'exit'
		}

		if response == "send message" {
			err := SendMessage(&server) // Call the SendMessage function to send a message
			if err != nil {
				fmt.Println("Failed to send message, please try again.")
				continue // Skip to the next iteration if sending message fails
			}
		}

		if response == "filter user" {
			fmt.Print("Enter the username to filter messages by: ")
			response, _ := reader.ReadString('\n')
			response = strings.TrimSpace(response) // Trim any extra whitespace or newline characters
			server.FilterByUser(response)          // Call the FilterByUser function to filter messages by user
			continue
		}

		if response == "filter message" {
			fmt.Print("Enter the keyword to filter by: ")
			response, _ := reader.ReadString('\n')
			response = strings.TrimSpace(response) // Trim any extra whitespace or newline characters
			server.FilterByKeyword(response)       // Call the FilterByUser function to filter messages by user
			continue
		}
	}
}

func CreateUser(users map[string]*User, id string) (*User, error) {
	// Helper function to create a new user
	if _, exists := users[id]; exists {
		return nil, fmt.Errorf("user with ID %s already exists", id)
	}
	inputChannel := make(chan Message)
	user, err := AddUser(users, id, inputChannel) // Ensure the user is added to the users map
	if err != nil {
		return nil, err // Return error if user could not be added
	}
	user.StartRouting() // Start routing messages for the new user
	return user, nil    // Return the newly created user
}

func AskUserForInput() string {
	// Helper function to ask for user input
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("What would you like to do? (send message / add user / filter user / filter message) [Type 'exit' to quit]: ")
	response, _ := reader.ReadString('\n')
	if strings.ToLower(strings.TrimSpace(response)) != "add user" && strings.ToLower(strings.TrimSpace(response)) != "exit" && strings.ToLower(strings.TrimSpace(response)) != "send message" && strings.ToLower(strings.TrimSpace(response)) != "filter user" && strings.ToLower(strings.TrimSpace(response)) != "filter message" {
		fmt.Println("Invalid response, please enter send message / add user / filter user / filter message")
		response = AskUserForInput() // Recursively ask for input again
	}
	return strings.ToLower(strings.TrimSpace(response))
}

func SendMessage(server *ChatServer) error {

	reader := bufio.NewReader(os.Stdin)
	var users = server.Users // Access the users map from the server

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
		fmt.Println("Please create the user first or check the username.")
		return fmt.Errorf("one or both users do not exist") // Return an error if either user does not exist
	}
	if !toExists {
		fmt.Printf("Error: User %s does not exist\n", userTo)
		fmt.Println("Please create the user first or check the username.")
		return fmt.Errorf("one or both users do not exist") // Return an error if either user does not exist
	}
	var newMessage Message
	newMessage.UserIDFrom = userFrom
	newMessage.UserIDTo = userTo
	newMessage.Content = message
	userToSendTo := users[userTo]                                                     // Get the recipient user object
	userSendFrom := users[userFrom]                                                   // Get the sender user object
	newMessage.Time = time.Now()                                                      // Set the time for the message before storing it in the history.
	userToSendTo.ReceivedMessages = append(userToSendTo.ReceivedMessages, newMessage) // Store the message in the recipient's message history
	userSendFrom.SentMessages = append(userSendFrom.SentMessages, newMessage)         // Store the message in the recipient's message history
	server.Messages = append(server.Messages, newMessage)                             // Store the message in the server's message history
	userToSendTo.Input <- newMessage                                                  // Send the message to the recipient's input channel

	newMessage.Time = time.Now() // Set the current time for the message
	time.Sleep(500 * time.Millisecond)
	return nil // Return nil to indicate success
}
