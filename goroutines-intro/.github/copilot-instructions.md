# Copilot Instructions for This Project

## Overview
This repository contains multiple Go projects, each located in its own numbered directory (e.g., `1/`, `2/`, etc.). Each directory includes a `main.go` file, which likely serves as the entry point for a standalone Go application or example.

## Key Conventions
- **Directory Structure**: Each numbered folder (e.g., `1/`, `2/`) represents an isolated Go project. These projects do not appear to share dependencies or code.
- **Entry Points**: The `main.go` file in each directory is the primary entry point for that project.
- **Go Modules**: If Go modules are used, ensure that each project is initialized with its own `go.mod` file. If not, the projects may rely on the GOPATH workspace.

## Developer Workflows

### Building and Running
To build and run a specific project, navigate to its directory and use the following commands:

```bash
cd <directory> # Replace <directory> with the folder name, e.g., 1, 2, etc.
go run main.go
```

### Testing
If tests are added to the projects, you can run them using:

```bash
go test ./...
```

Run this command from within the specific project directory.

### Debugging
Use the `dlv` (Delve) debugger for Go to debug a specific project:

```bash
dlv debug main.go
```

### Dependencies
If a project uses external dependencies, initialize a Go module and manage dependencies with:

```bash
go mod init <module-name>
go get <dependency>
```

## Project-Specific Patterns
- **Standalone Projects**: Each directory is self-contained. Avoid cross-referencing files or directories unless explicitly required.
- **Minimal Dependencies**: Projects appear to be lightweight and focused on Go's standard library. Use external libraries sparingly.

## Recommendations for AI Agents
- **Focus on Context**: When editing or generating code, ensure changes are scoped to the specific directory being worked on.
- **Follow Go Best Practices**: Adhere to idiomatic Go patterns, such as proper error handling and the use of `defer` for resource cleanup.
- **Avoid Assumptions**: Do not assume shared state or dependencies between directories.

## Example
To add a new feature to the project in directory `1/`, modify `1/main.go` and ensure the changes are isolated to that directory. For instance, to add a new function:

```go
// main.go
package main

import "fmt"

func newFeature() {
    fmt.Println("This is a new feature")
}

func main() {
    newFeature()
}
```

---

If you have any questions or need clarification, feel free to ask!