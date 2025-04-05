# 📦 Chat Application - Rust & Go Versions

This project is part of the **MSCS-632 Advanced Programming Languages Residency**. It showcases a CLI-based Chat Application implemented in **two different languages**: Rust and Go.

The goal is to compare how language-specific features handle the same functionality: user interaction, message routing, and concurrency.

---

## 🚀 Project Structure

```
MSCS-632-Residency-Project/
├── rust_chat_app/         # Rust implementation (turn-based chat)
├── go_chat_app/           # Go implementation (menu-driven chat)
└── README_chat_app.md     # You're here
```

---

## 🛠️ Installation Instructions

### 🔧 Install Rust
Visit the official [Rust installation page](https://www.rust-lang.org/tools/install) and run:
```bash
curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
```
> After installation, restart your terminal and verify with:
```bash
rustc --version
cargo --version
```

### 🔧 Install Go
Visit the official [Go installation page](https://go.dev/dl/) and download the latest stable release.
> On Ubuntu, you can use:
```bash
sudo apt install golang-go
```
Verify with:
```bash
go version
```

---

## 🦀 Rust Version

### ▶️ How to Run
```bash
cd rust_chat_app
cargo run
```

### 💬 Features
- Turn-based CLI chat
- Multiple users with custom names
- Exit per user after message
- Chat history displayed after all users exit
- Filtering by user and keyword



## 💬 Go Version

### ▶️ How to Run
```bash
cd go_chat_app
go run main.go chat.go user.go
```

### 💬 Features
- Menu-driven CLI system
- Dynamic user creation
- Message sending via sender/receiver prompt
- Real-time delivery using goroutines and channels
- Message filtering by sender and keyword



## 📚 Notes
- Both apps are **CLI-based** and designed for learning purposes.
- Concurrency is achieved via **async tokio** (Rust) and **goroutines/channels** (Go).
- Language-specific implementations were prioritized over UI/UX enhancements.

---

## 👥 Team Members
- Shabnam Shaikh
- Sakchham Sangroula
- Sandesh Pokharel
- Nihar Turumelle
- Romika Souda

---

## 🔗 GitHub Repo
[https://github.com/sanspokharel26677/MSCS-632-Residency-Project](https://github.com/sanspokharel26677/MSCS-632-Residency-Project)

