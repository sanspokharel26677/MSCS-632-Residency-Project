use std::collections::{HashMap, VecDeque};
use std::io::{self, Write};
use chrono::Local;

//////////////////////////////////////////////////////////////
// Welcome to the Turn-Based Rust Chat App!                //
// This app allows multiple users to take turns chatting.  //
// Each message is stored with user info and a timestamp.  //
//////////////////////////////////////////////////////////////

// Struct to represent a message
struct Message {
    user_id: String,
    content: String,
    timestamp: String,
}

// Chat server struct to store messages and stats
struct ChatServer {
    messages: VecDeque<Message>,
    message_count: HashMap<String, usize>,
    active_users: Vec<String>,
}

impl ChatServer {
    fn new(users: Vec<String>) -> Self {
        ChatServer {
            messages: VecDeque::new(),
            message_count: HashMap::new(),
            active_users: users,
        }
    }

    fn add_message(&mut self, msg: Message) {
        println!("[{}][{}]: {}", msg.timestamp, msg.user_id, msg.content);
        *self.message_count.entry(msg.user_id.clone()).or_insert(0) += 1;
        self.messages.push_back(msg);
    }

    fn show_summary(&self) {
        println!("\n=======================");
        println!("📜 Chat History Summary");
        println!("=======================");
        for msg in &self.messages {
            println!("[{}][{}]: {}", msg.timestamp, msg.user_id, msg.content);
        }

        println!("\n=======================");
        println!("📈 Messages Sent Per User");
        println!("=======================");
        for (user, count) in &self.message_count {
            println!("{}: {} message(s)", user, count);
        }

        println!("\n💡 Total Messages: {}", self.messages.len());
        println!("👥 Users Participated: {}", self.message_count.len());
        println!("✅ Chat Ended Successfully.");
    }
}

fn read_input(prompt: &str) -> String {
    print!("{}", prompt);
    io::stdout().flush().unwrap();
    let mut temp = String::new();
    io::stdin().read_line(&mut temp).unwrap();
    temp.trim().to_string()
}

fn main() {
    // 1. Welcome message
    println!("===============================");
    println!("👋 Welcome to the Rust Chat App!");
    println!("===============================");

    // 2. Ask for number of users
    let num_users: usize = loop {
        let input = read_input("How many users would you like to create? ");
        match input.parse() {
            Ok(n) if n > 0 => break n,
            _ => println!("Please enter a valid number greater than 0."),
        }
    };

    // 3. Ask for usernames
    let mut users = Vec::new();
    for i in 1..=num_users {
        let name = read_input(&format!("Enter username for User {}: ", i));
        users.push(name);
    }

    let mut server = ChatServer::new(users.clone());
    let mut exited_users = Vec::new();

    println!("\n🚀 Chat is starting! Type your messages one by one.");
    println!("Type `/exit` to leave the chat as a user.\n");

    // 4. Main loop: keep going while there are active users
    while !server.active_users.is_empty() {
        println!("\n👥 Active users: {:?}", server.active_users);
        let current = read_input("Who wants to send a message? ");

        if !server.active_users.contains(&current) {
            println!("⚠️ User '{}' is not active or doesn't exist.", current);
            continue;
        }

        let message = read_input(&format!("[{}] Enter your message: ", current));
        if message == "/exit" {
            println!("👋 {} has left the chat.", current);
            server.active_users.retain(|u| u != &current);
            exited_users.push(current.clone());
            continue;
        }

        let timestamp = Local::now().format("%H:%M:%S").to_string();
        let msg = Message {
            user_id: current.clone(),
            content: message,
            timestamp,
        };
        server.add_message(msg);
    }

    // 7. After all users exit
    println!("\n🎉 All users have left the chat.");
    server.show_summary();
}
