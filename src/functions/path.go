package functions

import (
	"strings"
)

func FormatPath(path string) string {
	return strings.ReplaceAll(path, "-", "/")
}
