package db

type Config struct {
	Host         string
	Port         int
	User         string
	Password     string
	DBName       string
	SSLMode      string
	ConnLifeTime int
	MaxIdleConns int
	MaxOpenConns int
	LogLevel     int
}
