# SSH Portfolio

[![Go](https://img.shields.io/badge/Language-Go-blue?logo=go&logoColor=white&label=Go%20v1.20)](https://golang.org) [![Bubbletea](https://img.shields.io/badge/Framework-Bubbletea-green)](https://github.com/charmbracelet/bubbletea) [![Wish](https://img.shields.io/badge/Framework-Wish-orange)](https://github.com/charmbracelet/wish)

A SSH server built with Go (Golang) and Wish to serve as a personal portfolio that showcases projects and skills.

## Project Goals

- Showcase projects and skills effectively.
- Ensure the server is robust and user-friendly.

## Getting Started

1. Ensure you have Go installed on your system.
2. Run the following command to execute the program:

   ```bash
   go run main.go
   ```

## Running the Application

1. Ensure you have Go installed on your system.
2. Generate an SSH key pair if you don't already have one. You can do this by running:

   ```bash
   ssh-keygen -t ed25519 -f .ssh/id_ed25519
   ```

3. Start the SSH server by running:

   ```bash
   go run main.go
   ```

4. Connect to the SSH server using an SSH client. For example:

   ```bash
   ssh -p 42069 localhost
   ```

   Replace `localhost` with the server's address if running on a remote machine.

## Build Configuration

This project uses `.air.toml` for build automation and configuration. The `.air.toml` file includes settings for:

- Build commands and output paths.
- File and directory exclusions during the build process.
- Logging and screen settings for build operations.

To build the project, you can use the following command:

```bash
# Build the project using the configuration in .air.toml
go build -o ./tmp/main .
```

## Project Structure

- `README.md`: Project documentation.
- `main.go`: Entry point of the application.
- `go.mod` and `go.sum`: Dependency management files.
- `pkg/`: Directory for reusable packages and libraries.
  - `tui/`: Contains the TUI (Text User Interface) implementation.
    - `model.go`: Core logic for the TUI.
- `tests/`: Directory for test files.
  - `tui_test.go`: Tests for the TUI package.
- `.github/`: Directory for GitHub-specific configurations.
  - `workflows/`: Contains GitHub Actions workflow files.

## License

This project is licensed under the MIT License.
