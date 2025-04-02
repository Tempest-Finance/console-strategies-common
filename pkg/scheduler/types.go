package scheduler

type Config struct {
	InitAddress  []string
	SelectDB     int
	Username     string
	Password     string
	DisableCache bool
}
