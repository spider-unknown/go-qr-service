package configs

// Server ...
type Server struct {
	Name    string `json:"name"`
	Version string `json:"-"`
	Addr    string `json:"addr" default:":7070"`
}

// Config ...
type Config struct {
	Server *Server `json:"server"`
}

// NewConfig ...
func NewConfig(name, version string, addr string) *Config {
	return &Config{
		Server: &Server{
			Name:    name,
			Version: version,
			Addr:    addr,
		},
	}
}
