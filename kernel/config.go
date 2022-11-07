package kernel

import "github.com/yrzs/k3cloud/object"

type Config struct {
	*object.Collection
}

func NewConfig(items *object.HashMap) *Config {

	return &Config{
		object.NewCollection(items),
	}
}
