package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/sluipmoord/sshportfolio/pkg/config"
	"github.com/sluipmoord/sshportfolio/pkg/tui"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/activeterm"
	"github.com/charmbracelet/wish/bubbletea"
	"github.com/charmbracelet/wish/logging"
)

type sshOutput struct {
	ssh.Session
	tty *os.File
}

func (s *sshOutput) Write(p []byte) (int, error) {
	return s.Session.Write(p)
}

func (s *sshOutput) Read(p []byte) (int, error) {
	return s.Session.Read(p)
}

func (s *sshOutput) Fd() uintptr {
	return s.tty.Fd()
}

func teaHandler(s ssh.Session) (tea.Model, []tea.ProgramOption) {
	pty, _, _ := s.Pty()
	sessionBridge := &sshOutput{
		Session: s,
		tty:     pty.Slave,
	}
	renderer := bubbletea.MakeRenderer(sessionBridge)
	command := s.Command()

	// Get client IP address from the SSH session
	clientAddr := s.RemoteAddr().String()
	host, _, _ := net.SplitHostPort(clientAddr)
	slog.Info("client connected", "ip", host)

	model, err := tui.NewModel(renderer, &host, command)
	if err != nil {
		return nil, []tea.ProgramOption{}
	}
	return model, []tea.ProgramOption{tea.WithAltScreen()}
}

func main() {
	// Load configuration
	cfg := config.Load()

	// Setup logger with configured log level
	if err := config.SetupLogger(cfg); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to set up logger: %v\n", err)
		os.Exit(1)
	}

	// Log startup information
	slog.Info("starting ssh server",
		"host", cfg.Host,
		"port", cfg.Port,
		"logLevel", cfg.LogLevel,
		"allowedClients", cfg.AllowedClients)

	// Create middleware stack
	middleware := []wish.Middleware{
		logging.Middleware(),
		activeterm.Middleware(),
		bubbletea.Middleware(teaHandler),
	}

	// Add IP filtering middleware if allowed clients are specified
	if len(cfg.AllowedClients) > 0 {
		middleware = append([]wish.Middleware{
			func(next ssh.Handler) ssh.Handler {
				return func(s ssh.Session) {
					clientAddr := s.RemoteAddr().String()
					host, _, _ := net.SplitHostPort(clientAddr)

					allowed := false
					for _, allowedIP := range cfg.AllowedClients {
						if host == allowedIP {
							allowed = true
							break
						}
					}

					if !allowed {
						slog.Warn("rejected connection from unauthorized client", "ip", host)
						s.Close()
						return
					}

					next(s)
				}
			},
		}, middleware...)
	}

	server, err := wish.NewServer(
		wish.WithAddress(net.JoinHostPort(cfg.Host, cfg.Port)),
		wish.WithHostKeyPath(cfg.HostKeyPath),
		wish.WithMiddleware(middleware...),
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
		slog.Info("Shutting down server...")
		if err := server.Shutdown(context.Background()); err != nil {
			slog.Error("Failed to shut down server", "error", err)
			os.Exit(1)
		}
		done <- true
	}()

	slog.Info(fmt.Sprintf("Starting SSH server on %s:%s", cfg.Host, cfg.Port))
	if err := server.ListenAndServe(); err != nil && err != ssh.ErrServerClosed {
		slog.Error("Failed to start SSH server", "error", err)
		os.Exit(1)
	}

	<-done
	slog.Info("Server stopped")
}
