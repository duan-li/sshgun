package bullet

import (
	"log"
	"os"
)

func WriteBullet(file string, data []byte) (bool, error) {
	dir := CurrentDir()

	err := os.WriteFile(dir+"/"+file, data, 0644)

	if err != nil {
		return false, err
	}

	return true, nil
}

func ReadBullet(file string) ([]byte, error) {

	dir := CurrentDir()

	data, err := os.ReadFile(dir + "/" + file)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func CurrentDir() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return dir
}
