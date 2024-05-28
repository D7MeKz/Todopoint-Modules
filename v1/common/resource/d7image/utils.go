package d7image

import (
	"github.com/google/uuid"
	"os"
	"path/filepath"
)

func FilenameGenerator() string {
	return uuid.NewString()
}

func NewFilePath(dirname string) (string, error) {
	dir := filepath.Join("./uploaded", dirname)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return "", err
		}
	}
	return dir, nil
}
