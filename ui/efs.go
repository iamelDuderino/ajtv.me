package ui

import "embed"

//go:embed "views" "static" "templates" "wasm"
var EFS embed.FS

//go:embed "styles.css"
var StyleSheet string
