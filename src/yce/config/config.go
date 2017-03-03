package config

type Config struct {
	// BuntDB
	// MySQL
	// Redis
}


func New() *Config {
	c := &Config {

	}
	return c
}

func (c *Config) ReadOsEnv() *Config {
	// redis endpoints
	// mysql endpoints
	// buntdb endpoints
	// siAgent endpoints

	return c
}


func (c *Config) Load() {

}
