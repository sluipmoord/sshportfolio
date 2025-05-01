package embedded

import (
	"embed"
)

//go:embed *.svg
var Assets embed.FS
