package util

import (
	"fmt"
	"strings"
)

// ParseHost IPV6改造
func ParseHost(host string) string {
	if strings.Contains(host, ":") {
		return fmt.Sprintf("[%s]", host)
	}
	return host
}
