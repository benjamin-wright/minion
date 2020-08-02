package loading

import (
	"fmt"
	"os"
	"strings"
)

func LoadVersion() (string, error) {
	file, err := os.Open("/input/version.txt")
	if err != nil {
		return "", fmt.Errorf("Failed to open versions file: %+v", err)
	}

	info, err := file.Stat()
	if err != nil {
		return "", fmt.Errorf("Failed to get file stats: %+v", err)
	}

	length := info.Size()
	bytes := make([]byte, length)

	_, err = file.Read(bytes)
	if err != nil {
		return "", fmt.Errorf("Failed to read file contents: %+v", err)
	}

	return strings.TrimSuffix(string(bytes), "\n"), nil
}
