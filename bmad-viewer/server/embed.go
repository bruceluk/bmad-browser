package main

import "embed"

//go:embed all:dist
var webDistFS embed.FS
