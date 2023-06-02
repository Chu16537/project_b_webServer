package conf

type Config struct {
	WebServer `yaml:"webserver"`
}

type WebServer struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}
