package model

//Configuration properties
type Configuration struct {
	Server   ServerConfig
	Database DBConfig
}

//ServerConfig properties
type ServerConfig struct {
	Port string
}

//DBConfig properties
type DBConfig struct {
	DBName     string
	DBUser     string
	DBPassword string
	DBServer   string
	DBPort     string
}
