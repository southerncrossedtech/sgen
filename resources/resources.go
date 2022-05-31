package resources

import "embed"

// Embedded contains the resources we want to be bundled into our binary at compile time.

//go:embed templates
var Templates embed.FS
