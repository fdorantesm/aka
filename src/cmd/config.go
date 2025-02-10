package cmd

type Config struct {
	Version string
}

var compiledVersion string = "0.0.0"

func LoadConfig() Config {
	return Config{Version: compiledVersion}
}
