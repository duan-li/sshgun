package bullet

import (
	"gopkg.in/yaml.v3"
)

type Bullet struct {
	Servers  []server `yaml:"servers"`
	Commands []string `yaml:"commands"`
}

type server struct {
	Ip         string
	Port       int
	Username   string
	Password   string
	Supassword string
}

func Paser(content []uint8) (map[string]Bullet, error) {
	var data map[string]Bullet

	err := yaml.Unmarshal(content, &data)

	if err != nil {
		return nil, err
	}

	return data, nil
}
