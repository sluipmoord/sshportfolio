package main

import (
	"log"
	"net"
	"sshportfolio/pkg/tui"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/logging"
)

const (
	host = "localhost"
	port = "42069"
)

func portfolioMiddleware(projects []string) wish.Middleware {
	return func(handler ssh.Handler) ssh.Handler {
		return func(s ssh.Session) {
			p := tea.NewProgram(tui.Model{Projects: projects}, tea.WithInput(s), tea.WithOutput(s))
			if _, err := p.Run(); err != nil {
				log.Printf("Error starting TUI: %v", err)
			}
			handler(s)
		}
	}
}

func main() {
	projects := []string{
		"Project 1: SSH Server",
		"Project 2: TUI Portfolio",
		"Project 3: Go Web App",
	}

	server, err := wish.NewServer(
		wish.WithAddress(net.JoinHostPort(host, port)),
		wish.WithHostKeyPath(".ssh/id_ed25519"),
		wish.WithMiddleware(
			logging.Middleware(),
			portfolioMiddleware(projects),
		),
	)
	if err != nil {
		log.Fatalf("Failed to create SSH server: %v", err)
	}

	log.Println("Starting SSH server on :42069")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start SSH server: %v", err)
	}
}
