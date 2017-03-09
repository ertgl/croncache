package v1

type Config struct {
	Host               string
	Port               int
	Database           string
	Password           string
	ConnectionPoolSize int
}

func NewConfig() *Config {
	c := &Config{
		Host:               "localhost",
		Port:               6379,
		Database:           "1",
		Password:           "",
		ConnectionPoolSize: 5,
	}
	return c
}
