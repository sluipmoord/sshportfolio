# SSH Portfolio

[![Go](https://img.shields.io/badge/Language-Go-blue?logo=go&logoColor=white&label=Go%20v1.23)](https://golang.org) [![Bubbletea](https://img.shields.io/badge/Framework-Bubbletea-green)](https://github.com/charmbracelet/bubbletea) [![Wish](https://img.shields.io/badge/Framework-Wish-orange)](https://github.com/charmbracelet/wish)

A SSH server built with Go (Golang) and Wish to serve as a personal portfolio that showcases projects and skills. This project provides both a SSH server and a standalone CLI application, both using the same TUI (Terminal User Interface) components.

## Project Goals

- Showcase projects and skills effectively.
- Ensure the server is robust and user-friendly.

## Getting Started

1. Ensure you have Go installed on your system (version 1.23 or later recommended).
2. Clone this repository:

   ```bash
   git clone https://github.com/yourusername/sshportfolio.git
   cd sshportfolio
   ```

3. Install dependencies:

   ```bash
   go mod download
   ```

## Running the Application

### CLI Mode (Local)

For testing and development, you can run the application as a local CLI without SSH:

```bash
go run ./cmd/cli/main.go
```

### SSH Server Mode

1. Generate an SSH key pair if you don't already have one:

   ```bash
   mkdir -p .ssh
   ssh-keygen -t ed25519 -f .ssh/id_ed25519
   ```

2. Start the SSH server:

   ```bash
   go run ./cmd/ssh/main.go
   ```

3. Connect to the SSH server using any SSH client:

   ```bash
   ssh -p 42069 localhost
   ```

## Development with Air

This project supports hot-reloading using [Air](https://github.com/cosmtrek/air). To use it:

1. Install Air:

   ```bash
   go install github.com/cosmtrek/air@latest
   ```

2. Run the SSH server with hot reloading:

   ```bash
   air
   ```

The Air configuration in `.air.toml` is set up to rebuild and restart the SSH server when code changes are detected.

## Project Architecture

### Components

- **SSH Server** (`cmd/ssh/main.go`): Provides portfolio access via SSH
- **CLI Application** (`cmd/cli/main.go`): Local terminal interface for the portfolio
- **TUI Package** (`pkg/tui/`): Shared terminal UI components used by both applications
- **Theme Package** (`pkg/tui/theme/`): Styling components for consistent visual appearance

### Project Structure

- `README.md`: Project documentation
- `go.mod` and `go.sum`: Dependency management files
- `cmd/`: Command applications
  - `cli/`: CLI application entry point
  - `ssh/`: SSH server entry point
- `pkg/`: Reusable packages
  - `tui/`: Terminal UI components
    - `root.go`: Main TUI model and view logic
    - `theme/`: UI styling components
      - `theme.go`: Theme definition and style utilities
      - `huh.go`: Form styling utilities
- `tests/`: Test files
- `.air.toml`: Configuration for the Air hot-reload tool

## Dependencies

This project uses several libraries from the [Charm](https://charm.sh) ecosystem:

- [Bubbletea](https://github.com/charmbracelet/bubbletea): TUI framework
- [Lipgloss](https://github.com/charmbracelet/lipgloss): Style definitions for terminal applications
- [Wish](https://github.com/charmbracelet/wish): SSH server framework
- [Huh](https://github.com/charmbracelet/huh): Form/input components

## License

This project is licensed under the MIT License.
