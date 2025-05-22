// Go (Golang) is a statically typed, compiled programming language created by Google.
// It is designed for simplicity, performance, and built-in support for concurrency.
// Go is often used for backend development, cloud services, command-line tools, and scalable systems.

// Why use Go?
// - Fast compilation and execution
// - Clean and minimal syntax
// - Powerful concurrency model (goroutines and channels)
// - Strong standard library
// - Cross-platform support (compile for any OS)
// - Excellent tooling (formatting, testing, documentation)

// package main
// Every Go file starts with a package declaration.
// The 'main' package is special: it defines a standalone executable application.
// The Go runtime looks for a package named 'main' and runs its 'main' function as the entry point.

// import "fmt"
// The import statement brings in external or standard packages.
// 'fmt' is a standard Go package for formatted I/O (input/output).
// It includes functions like Print, Println, Scanf, etc.

// func main() { ... }
// 'func' is used to define a function in Go.
// The 'main' function is the entry point of every Go program.
// Execution starts from this function when you run the application.

package main

import "fmt"

func main() {
	fmt.Print("Hello World!")
}

// To run this program:
// go run app.go

// To compile it into a binary:
// go build
// ./first-app (on Unix/Mac) or first-app.exe (on Windows)

// Go Modules:
// A Go module is a collection of Go packages stored in a file tree with a `go.mod` file at its root.
// It defines the module path (project name) and manages dependencies (versions of other modules).
// Modules are the standard way to manage project dependencies and build reproducible builds in modern Go.

// To initialize a module:
// Run `go mod init <module-name>` in the root of your project.
// Example: `go mod init example.com/first-app`

// This creates a go.mod file, which contains:
// - The module path
// - The Go version used
// - Any dependencies your project imports (added automatically by `go get` or on build)

// To view your current module info:
// Check the `go.mod` file
// You may also have a `go.sum` file, which ensures dependency integrity with checksums
