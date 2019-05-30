package utils

import (
	"os"
	"regexp"
	"strings"
)

func ExistsPath(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func BuildPathIfNotExist(path string) error {
	ok, err := ExistsPath(path)
	if !ok {
		return os.MkdirAll(path, 0644)
	}
	return err
}

func TrimInvalidFilePathChars(path string) string {
	path = strings.TrimSpace(path)
	re := regexp.MustCompile("[\\\\/:*?\"<>|]")
	return re.ReplaceAllString(path, "")
}
