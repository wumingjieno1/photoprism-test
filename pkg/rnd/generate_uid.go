package rnd

import (
	"strconv"
	"time"
)

// GenerateUID returns a unique id with prefix as string.
func GenerateUID(prefix byte) string {
	result := make([]byte, 0, 16)
	result = append(result, prefix)
	result = append(result, strconv.FormatInt(time.Now().UTC().Unix(), 36)[0:6]...)
	result = append(result, GenerateToken(9)...)

	return string(result)
}
