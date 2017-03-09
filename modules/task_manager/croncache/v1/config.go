package v1

type Config struct {
	Tasks map[string]string
}

func NewConfig() *Config {
	c := &Config{
		Tasks: make(map[string]string, 0),
	}
	return c
}
