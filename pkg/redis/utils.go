package redis

import (
	"strings"
)

func FormatKey(args ...string) string {
	return strings.Join(args, ":")
}
