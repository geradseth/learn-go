### Learning Go 🚀

A structured repository dedicated to mastering the Go (Golang) programming language. This project documents my journey from learning basic syntax to understanding advanced Go idioms, concurrency models, and microservices. 

### 📌 Project Overview

The goal of this repository is to build a solid foundation in Go by implementing core concepts, algorithm practices, and mini-applications. It acts as both a personal code sandbox and a quick-reference guide for Go design patterns and best practices.
 

### 🛠️ Prerequisites

To run the code in this repository, you need to have the following installed: 

* **Go**: Version 1.21 or higher
* A code editor like **Visual Studio Code** (with the official Go extension) or **GoLand**

### 📂 Repository Structure

The project is organized by topic complexity to make navigation intuitive: 

text

.
├── 01-basics/            # Basic syntax, variables, and control flow
├── 02-data-structures/   # Arrays, slices, maps, and custom data structures
├── 03-concurrency/       # Goroutines, channels, and sync packages
├── 04-web-development/   # Standard library HTTP servers and REST APIs
├── 05-testing/           # Unit tests, benchmarks, and mocks
├── go.mod                # Go module definition
└── README.md             # Project documentation

Use code with caution.

### 🚀 Getting Started

1. **Clone the repository:** 

bash

git clone https://github.com/YOUR_USERNAME/learning-go.git
cd learning-go

Use code with caution.
2. **Initialize or verify modules:** 

bash

go mod tidy

Use code with caution.
3. **Navigate to a specific topic and run the code:** 

bash

cd 01-basics
go run main.go

Use code with caution.

### 🧠 Core Concepts Explored

ModuleTopics CoveredKey Implementation
****01-Basics****
Variables, Types, Pointers, Loops, FunctionsControl structures and standard I/O
****02-Data Structures****
Slices, Maps, Structs, InterfacesCustom interface implementations
****03-Concurrency****
Goroutines, Channels, Select, MutexesWorker pool patterns
****04-Web Dev****
net/http, Routing, JSON handlingA lightweight CRUD REST API
****05-Testing****
Unit tests, Table-driven tests, Benchmarkingtesting package deep-dive

### 🧪 Running Tests

Go has an excellent built-in testing framework. You can run all tests across the repository using: 

bash

go test ./... -v

Use code with caution.

To run benchmarks: 

bash

go test ./... -bench=.

Use code with caution.

### 📚 Useful Resources

* [A Tour of Go](https://go.dev/tour/) - The official interactive introduction to Go.
* [Effective Go](https://go.dev/doc/effective_go) - Tips for writing clear, idiomatic Go code.
* [Go by Example](https://gobyexample.com/) - A hands-on introduction to Go using annotated program examples.

### 📄 License

This repository is configured under the [MIT License](LICENSE). Feel free to use the code for your own learning purposes.
