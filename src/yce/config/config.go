package config

type Config struct {
	// BuntDB
	// MySQL
	// Redis
}


func New() *Config {
	c := &Config {

	}
	c.readOsEnv()
}

func (c *Config) readOsEnv() {
	// redis endpoints
	// mysql endpoints
	// buntdb endpoints
	// siAgent endpoints
}


func (c *Config) Load() {

}
