package utils

import (
	"fmt"
)

func BinaryNotFoundError(binary string) error {
	return fmt.Errorf("%s not found", binary)
}
