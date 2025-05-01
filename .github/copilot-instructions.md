<!-- Use this file to provide workspace-specific custom instructions to Copilot. For more details, visit https://code.visualstudio.com/docs/copilot/copilot-customization#_use-a-githubcopilotinstructionsmd-file -->

# A SSH server built with Go (Golang) and Wish to serve as a personal portfolio that showcases projects and skills

use Go as the programming language
use Bubbletea as the TUI framework
use Wish as the SSH server framework

## Project Goals

- Showcase projects and skills effectively.
- Ensure the server is robust and user-friendly.

## Documentation Updates

- Always update the docs and README.md file with the latest changes

## Project Structure

- `main.go`: Entry point of the application.
- `.github/copilot-instructions.md`: Workspace-specific instructions for Copilot.
- `README.md`: Project documentation.
- `LICENSE`: Project license.
- `assets/`: Directory for static assets (images, styles, etc.).
- `cmd/`: Directory for command-line tools.
- `pkg/`: Directory for reusable packages and libraries.
- `internal/`: Directory for internal packages.
- `tests/`: Directory for test files.
- `docs/`: Directory for documentation files.
- `scripts/`: Directory for scripts and automation.

## Commit Message Generation

- Use conventional commit message format.
- Use imperative mood.
- Use present tense.
- Use lowercase letters.
- Use a short summary.
- Use a body to explain the changes.
- Use bullet points for multiple changes.
- Use a footer for breaking changes.
- Always the format

  ```bash
    <type>[optional scope]: <description> 
    [optional body] 
    [optional footer(s)]
  ```

## Code Style Guidelines

- **Imports**: Standard library first, followed by external dependencies, then local packages
- **Types**: Define types at the top of files, use custom structs for domain models
- **Naming**: PascalCase for exported identifiers, camelCase for private
- **Error Handling**: Explicit error returns with context, central error display mechanism
- **Modules**: Organized by logical domain (api, tui, resource)
- **Testing**: `_test.go` files with context-based testing
- **Documentation**: Use Go doc comments for public functions and types
- **UI Components**: Composition-based Bubble Tea components with Model-View pattern
- **Formatting**: Standard Go formatting with `go fmt`
  