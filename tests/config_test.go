package tests

import (
	"flag"
	"log/slog"
	"os"
	"testing"

	"github.com/sluipmoord/sshportfolio/pkg/config"
)

func TestParseLogLevel(t *testing.T) {
	tests := []struct {
		input    string
		expected slog.Level
	}{
		{"debug", slog.LevelDebug},
		{"info", slog.LevelInfo},
		{"warn", slog.LevelWarn},
		{"error", slog.LevelError},
		{"invalid", slog.LevelInfo}, // Default to info for invalid levels
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := config.ParseLogLevel(test.input)
			if result != test.expected {
				t.Errorf("ParseLogLevel(%s) = %v, expected %v", test.input, result, test.expected)
			}
		})
	}
}

func TestConfigLoading(t *testing.T) {
	// Save original os.Args
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	// Test default values
	os.Args = []string{"cmd"}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)

	// Clear environment variables that might affect the test
	os.Unsetenv("SSH_HOST")
	os.Unsetenv("SSH_PORT")
	os.Unsetenv("LOG_LEVEL")
	os.Unsetenv("SSH_ALLOWED_CLIENTS")

	cfg := config.Load()

	if cfg.Host != "localhost" {
		t.Errorf("Default host should be localhost, got %s", cfg.Host)
	}
	if cfg.Port != "42069" {
		t.Errorf("Default port should be 42069, got %s", cfg.Port)
	}
	if cfg.LogLevel != "info" {
		t.Errorf("Default log level should be info, got %s", cfg.LogLevel)
	}

	// Test with command-line args
	os.Args = []string{"cmd", "-host", "0.0.0.0", "-port", "2222", "-loglevel", "debug", "-allowed-clients", "192.168.1.1, 10.0.0.1"}

	// Reset flags
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)

	cfg = config.Load()

	if cfg.Host != "0.0.0.0" {
		t.Errorf("Host should be 0.0.0.0, got %s", cfg.Host)
	}
	if cfg.Port != "2222" {
		t.Errorf("Port should be 2222, got %s", cfg.Port)
	}
	if cfg.LogLevel != "debug" {
		t.Errorf("Log level should be debug, got %s", cfg.LogLevel)
	}
	if len(cfg.AllowedClients) != 2 {
		t.Errorf("Should have 2 allowed clients, got %d", len(cfg.AllowedClients))
	}
	if cfg.AllowedClients[0] != "192.168.1.1" {
		t.Errorf("First allowed client should be 192.168.1.1, got %s", cfg.AllowedClients[0])
	}

	// Test with environment variables
	os.Args = []string{"cmd"} // Reset args
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)

	os.Setenv("SSH_HOST", "127.0.0.1")
	os.Setenv("SSH_PORT", "9999")
	os.Setenv("LOG_LEVEL", "warn")

	cfg = config.Load()

	if cfg.Host != "127.0.0.1" {
		t.Errorf("Host should be 127.0.0.1, got %s", cfg.Host)
	}
	if cfg.Port != "9999" {
		t.Errorf("Port should be 9999, got %s", cfg.Port)
	}
	if cfg.LogLevel != "warn" {
		t.Errorf("Log level should be warn, got %s", cfg.LogLevel)
	}
}

func TestLoggerSetup(t *testing.T) {
	// Create a temp file for testing logger
	tempFile, err := os.CreateTemp("", "log-test")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())
	tempFile.Close()

	cfg := &config.Config{
		LogLevel:    "debug",
		LogFilePath: tempFile.Name(),
	}

	err = config.SetupLogger(cfg)
	if err != nil {
		t.Errorf("SetupLogger should not return error, got: %v", err)
	}

	// Test with stdout
	cfg.LogFilePath = "stdout"
	err = config.SetupLogger(cfg)
	if err != nil {
		t.Errorf("SetupLogger with stdout should not return error, got: %v", err)
	}

	// Test with stderr
	cfg.LogFilePath = "stderr"
	err = config.SetupLogger(cfg)
	if err != nil {
		t.Errorf("SetupLogger with stderr should not return error, got: %v", err)
	}
}
