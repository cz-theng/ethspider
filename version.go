package ethspider

import (
	"fmt"
)

const (
	major = 0
	minor = 1
	patch = 0
)

// Version return eth spider's version
func Version() string {
	return fmt.Sprintf("ethspider[%d.%d.%d]", major, minor, patch)
}
