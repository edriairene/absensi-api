package config

type Config struct {
	ListenURL string `toml:"listen-url"`
	ListenPort int `toml:"listen-port"`
	FormURL string `toml:"form-url"`
	FormPort int `toml:"form-port"`
	ReportURL string `toml:"report-url"`
	ReportPort int `toml:"report-port"`
}