# sshgleamthing

This is a clean Go (Golang) project.

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

## Project Structure

- `main.go`: Entry point of the application.
- `.github/copilot-instructions.md`: Workspace-specific instructions for Copilot.

## License

This project is licensed under the MIT License.