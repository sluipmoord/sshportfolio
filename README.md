# SSH Portfolio

[![Go](https://img.shields.io/badge/Language-Go-blue?logo=go&logoColor=white&label=Go%20v1.23)](https://golang.org) [![Bubbletea](https://img.shields.io/badge/Framework-Bubbletea-green)](https://github.com/charmbracelet/bubbletea) [![Wish](https://img.shields.io/badge/Framework-Wish-orange)](https://github.com/charmbracelet/wish)

A SSH server built with Go (Golang) and Wish to serve as a personal portfolio that showcases projects and skills. This project provides both a SSH server and a standalone CLI application, both using the same TUI (Terminal User Interface) components.

## Project Goals

- Showcase projects and skills effectively.
- Ensure the server is robust and user-friendly.
- Maintain high code quality with Go best practices.

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

Or use the Makefile:

```bash
make run/cli
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

   Or use the Makefile:

   ```bash
   make run
   ```

3. Connect to the SSH server using any SSH client:

   ```bash
   ssh -p 42069 localhost
   ```

## Configuration

Both the SSH server and CLI application share a common configuration system and can be configured using command-line flags or environment variables:

| Flag | Environment Variable | Default | Description |
|------|---------------------|---------|-------------|
| `-host` | `SSH_HOST` | `localhost` | Host address to bind the SSH server |
| `-port` | `SSH_PORT` | `42069` | Port to listen on |
| `-key` | `SSH_KEY_PATH` | `.ssh/id_ed25519` | Path to SSH host key |
| `-loglevel` | `LOG_LEVEL` | `info` | Log level: debug, info, warn, error |
| `-logfile` | `LOG_FILE` | `output.log` | Log file path (or "stdout"/"stderr") |
| `-allowed-clients` | `SSH_ALLOWED_CLIENTS` | (none) | Comma-separated list of allowed client IPs |

Examples:

```bash
# Run on all network interfaces with custom port and debug logging
go run ./cmd/ssh/main.go -host 0.0.0.0 -port 2222 -loglevel debug

# Log to standard error instead of a file
go run ./cmd/cli/main.go -logfile stderr -loglevel debug

# Run with IP restrictions (only allow specific clients)
go run ./cmd/ssh/main.go -allowed-clients 192.168.1.10,192.168.1.11

# Using environment variables
export SSH_HOST=0.0.0.0
export SSH_PORT=2222
export LOG_LEVEL=debug
export LOG_FILE=stdout
go run ./cmd/ssh/main.go
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

## Terminal UI Components

The application features a TUI built with Bubbletea and includes the following pages:

- **Menu Page**: Navigation hub for accessing all other sections
- **About Me**: Personal information and introduction
- **Projects**: Showcase of your development projects
- **Skills**: List of technical skills and competencies
- **README**: Interactive display of this README.md file with proper Markdown rendering

Navigation controls:

- Arrow keys or `j`/`k` to navigate menus
- `Enter` to select
- `Esc` to go back to menu
- `q` or `Ctrl+C` to quit
- For README and other content pages with scrolling:
  - Arrow keys, `Page Up`/`Page Down`, or mouse wheel to scroll

## Embedded Assets

The project includes an assets package (`pkg/assets`) that manages embedded resources using Go's embed directive. This allows you to:

1. Include static assets (like SVG icons, images, and style files) directly in the binary
2. Access these resources efficiently at runtime without external file dependencies

To add new assets:

1. Place the file in the `pkg/assets/embedded/` directory
2. The `assets.go` file automatically handles embedding these resources
3. Access them in your code through the assets package API

## Customizing Your Portfolio

To personalize this portfolio for your own use:

1. **Update the About Me Page**: Modify the content function in `pkg/tui/root.go` for the About Me page to include your personal information.

2. **Projects Section**: Edit the Projects page content to showcase your own work.

3. **Skills Display**: Customize the Skills page to reflect your technical capabilities.

4. **Theme Customization**: Modify the theme settings in `pkg/tui/theme/theme.go` to match your preferred color scheme and styling.

5. **Embedded Assets**: Replace or add custom icons and images in the `pkg/assets/embedded/` directory.

## Testing

The project includes a comprehensive test suite:

```bash
# Run all tests
make test
```

Test files are located in the `tests/` directory and cover:

- Configuration validation and loading
- TUI component rendering and behavior
- Integration tests for SSH server functionality

## Code Quality

The project follows Go best practices to ensure code quality:

- Custom types for context keys to prevent key collisions
- Consistent error handling and logging
- Comprehensive testing
- Clean separation of concerns with a well-defined package structure

## Project Architecture

### Components

- **SSH Server** (`cmd/ssh/main.go`): Provides portfolio access via SSH
- **CLI Application** (`cmd/cli/main.go`): Local terminal interface for the portfolio
- **TUI Package** (`pkg/tui/`): Shared terminal UI components used by both applications
- **Config Package** (`pkg/config/`): Shared configuration and logging setup
- **Theme Package** (`pkg/tui/theme/`): Styling components for consistent visual appearance
- **Assets Package** (`pkg/assets/`): Embedded resources management

### Project Structure

- `README.md`: Project documentation
- `go.mod` and `go.sum`: Dependency management files
- `.github/`: GitHub-specific files (e.g., workflows, issue templates)
- `cmd/`: Command applications
  - `cli/`: CLI application entry point
  - `ssh/`: SSH server entry point
- `pkg/`: Reusable packages
  - `assets/`: Embedded static assets
    - `assets.go`: Asset loader
    - `embedded/`: Static assets directory
      - `embedded.go`: Embedding declaration
      - `github.svg`: Example SVG icon
  - `config/`: Configuration and logging setup
    - `config.go`: Configuration loading and logger setup
  - `tui/`: Terminal UI components
    - `root.go`: Main TUI model and view logic
    - `footer.go`: Footer components
    - `header.go`: Header components
    - `menu.go`: Menu navigation components
    - `pages.go`: Content pages
    - `theme/`: UI styling components
      - `theme.go`: Theme definition and style utilities
      - `huh.go`: Form styling utilities
- `tests/`: Test files
  - `config_test.go`: Tests for configuration
  - `main_test.go`: Main test entry point
  - `tui_test.go`: Tests for TUI components
- `Makefile`: Build and development commands

## Makefile Commands

The project includes a Makefile with several useful commands:

```bash
make test         run all tests
make test/cover   run all tests and display coverage
make develop      run all tests and start the server
make cli          run the CLI
make build        build the server
make build/cli    build the CLI
make run          run the server
make run/cli      run the CLI
make clean        remove build artifacts
make tidy         format code and tidy modfile
make audit        run quality control checks
make help         print this help message
```

## Dependencies

This project uses several libraries from the [Charm](https://charm.sh) ecosystem:

- [Bubbletea](https://github.com/charmbracelet/bubbletea): TUI framework
- [Lipgloss](https://github.com/charmbracelet/lipgloss): Style definitions for terminal applications
- [Wish](https://github.com/charmbracelet/wish): SSH server framework
- [Huh](https://github.com/charmbracelet/huh): Form/input components
- [Glamour](https://github.com/charmbracelet/glamour): Markdown rendering

## License

This project is licensed under the MIT License.
