package embedded

import (
	"embed"
)

//go:embed *.svg *.md
var Assets embed.FS
