package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"sshportfolio/pkg/tui"
	"syscall"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/logging"
)

const (
	host = "localhost"
	port = "42069"
)

func portfolioMiddleware(pages []string) wish.Middleware {
	return func(handler ssh.Handler) ssh.Handler {
		return func(s ssh.Session) {
			p := tea.NewProgram(tui.Portfolio{Pages: pages}, tea.WithInput(s), tea.WithOutput(s))
			if _, err := p.Run(); err != nil {
				log.Printf("Error starting TUI: %v", err)
			}
			handler(s)
		}
	}
}

func main() {
	pages := []string{
		"Hello",
		"About",
		"Projects",
	}

	server, err := wish.NewServer(
		wish.WithAddress(net.JoinHostPort(host, port)),
		wish.WithHostKeyPath(".ssh/id_ed25519"),
		wish.WithMiddleware(
			logging.Middleware(),
			portfolioMiddleware(pages),
		),
	)
	if err != nil {
		log.Fatalf("Failed to create SSH server: %v", err)
	}

	// Channel to listen for termination signals
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	// Channel to signal server shutdown
	done := make(chan bool, 1)

	go func() {
		<-signalChan
		log.Println("Shutting down server...")
		if err := server.Shutdown(context.Background()); err != nil {
			log.Fatalf("Failed to shut down server: %v", err)
		}
		done <- true
	}()

	log.Println("Starting SSH server on :42069")
	if err := server.ListenAndServe(); err != nil && err != ssh.ErrServerClosed {
		log.Fatalf("Failed to start SSH server: %v", err)
	}

	<-done
	log.Println("Server stopped.")
}
