package boltbrowserweb

import (
	"embed"
)

//go:embed css/* fonts/* js/* html/*
var Assets embed.FS
