package config

import (
	"flag"
	"io"
	"log/slog"
	"os"
	"strings"
)

// Config holds all configuration options for both SSH server and CLI
type Config struct {
	// SSH Server specific options
	Host           string
	Port           string
	HostKeyPath    string
	AllowedClients []string

	// Common options
	LogLevel    string
	LogFilePath string
}

// getEnv retrieves an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// ParseLogLevel converts a string log level to slog.Level
func ParseLogLevel(level string) slog.Level {
	switch strings.ToLower(level) {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

// Load loads configuration from flags and environment variables
func Load() *Config {
	// Default values
	config := &Config{
		Host:           "localhost",
		Port:           "42069",
		HostKeyPath:    ".ssh/id_ed25519",
		LogLevel:       "info",
		LogFilePath:    "output.log",
		AllowedClients: []string{},
	}

	// Command line flags
	flag.StringVar(&config.Host, "host", getEnv("SSH_HOST", config.Host), "SSH server host (env: SSH_HOST)")
	flag.StringVar(&config.Port, "port", getEnv("SSH_PORT", config.Port), "SSH server port (env: SSH_PORT)")
	flag.StringVar(&config.HostKeyPath, "key", getEnv("SSH_KEY_PATH", config.HostKeyPath), "Path to SSH host key (env: SSH_KEY_PATH)")
	flag.StringVar(&config.LogLevel, "loglevel", getEnv("LOG_LEVEL", config.LogLevel), "Log level: debug, info, warn, error (env: LOG_LEVEL)")
	flag.StringVar(&config.LogFilePath, "logfile", getEnv("LOG_FILE", config.LogFilePath), "Log file path (env: LOG_FILE)")

	// Parse allowed client IPs (comma-separated list)
	allowedClientsFlag := flag.String("allowed-clients", getEnv("SSH_ALLOWED_CLIENTS", ""), "Comma-separated list of allowed client IPs (env: SSH_ALLOWED_CLIENTS)")

	flag.Parse()

	// Process allowed clients if provided
	if *allowedClientsFlag != "" {
		config.AllowedClients = strings.Split(*allowedClientsFlag, ",")
		for i, client := range config.AllowedClients {
			config.AllowedClients[i] = strings.TrimSpace(client)
		}
	}

	return config
}

// SetupLogger configures the global slog logger based on the provided configuration
func SetupLogger(cfg *Config) error {
	var writers []io.Writer

	// If log file path is "stdout" or "stderr", just use those
	if cfg.LogFilePath == "stdout" {
		writers = append(writers, os.Stdout)
	} else if cfg.LogFilePath == "stderr" {
		writers = append(writers, os.Stderr)
	} else {
		// Otherwise, open the specified log file
		logFile, err := os.OpenFile(cfg.LogFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			return err
		}

		// Add both the file and stdout to writers
		writers = append(writers, logFile, os.Stdout)
	}

	// Create a multi-writer that writes to all configured outputs
	multiWriter := io.MultiWriter(writers...)

	// Set up the logger with the configured level
	level := ParseLogLevel(cfg.LogLevel)
	handler := slog.NewTextHandler(multiWriter, &slog.HandlerOptions{Level: level})
	logger := slog.New(handler)
	slog.SetDefault(logger)

	return nil
}
