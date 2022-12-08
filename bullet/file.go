package bullet

import (
	"log"
	"os"
)

// WriteBullet write content to bullet file
func WriteBullet(file string, data []byte) (bool, error) {
	dir := CurrentDir()

	err := os.WriteFile(dir+"/"+file, data, 0644)

	if err != nil {
		return false, err
	}

	return true, nil
}

// ReadBullet read content from bullet file
func ReadBullet(file string) ([]byte, error) {

	dir := CurrentDir()

	data, err := os.ReadFile(dir + "/" + file)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// CurrentDir get current directory
func CurrentDir() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return dir
}
