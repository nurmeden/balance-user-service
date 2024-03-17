package utils

func GetConfigPath(configPath string) string {
	if configPath == "docker" {
		return "./config/config-docker"
	}

	return "./configs/.env"
}
