package assets

import (
	"fmt"
	"io/fs"
	"strings"

	"github.com/sluipmoord/sshportfolio/pkg/assets/embedded"
)

// GetAsset loads an asset from the embedded filesystem
func GetAsset(path string) ([]byte, error) {
	return embedded.Assets.ReadFile(path)
}

// GetSVG loads an SVG asset from the embedded filesystem
func GetSVG(name string) ([]byte, error) {
	if !strings.HasSuffix(name, ".svg") {
		name = fmt.Sprintf("%s.svg", name)
	}

	return GetAsset(name)
}

// ListAssets returns a list of all embedded assets
func ListAssets() ([]string, error) {
	var assets []string

	err := fs.WalkDir(embedded.Assets, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() {
			assets = append(assets, path)
		}

		return nil
	})

	return assets, err
}

// SVGToASCII converts a simple SVG to ASCII art
func SVGToASCII(svgName string, width, height int) (string, error) {
	// Read SVG file from embedded assets
	data, err := GetSVG(svgName)
	if err != nil {
		return "", fmt.Errorf("failed to read SVG file: %w", err)
	}

	// Check if the file contains GitHub SVG
	if strings.Contains(string(data), "GitHub</title>") {
		// Simple GitHub representation
		return "󰊤", nil // GitHub icon (if terminal supports it) or use " GH "
	}

	// For other SVGs, return a generic representation
	return "[icon]", nil
}
