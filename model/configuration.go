package model

//Configuration properties
type Configuration struct {
	Server ServerConfig
}

//ServerConfig properties
type ServerConfig struct {
	Port int
}
