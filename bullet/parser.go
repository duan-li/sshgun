package bullet

import (
	"gopkg.in/yaml.v3"
)

// Bullet is a struct for bullet
type Bullet struct {
	Servers  []server `yaml:"servers"`
	Commands []string `yaml:"commands"`
}

// server is a struct for server
type server struct {
	Ip         string
	Port       int
	Username   string
	Password   string
	Supassword string
}

// Paser parse yaml file
func Paser(content []uint8) (map[string]Bullet, error) {
	var data map[string]Bullet

	err := yaml.Unmarshal(content, &data)

	if err != nil {
		return nil, err
	}

	return data, nil
}
