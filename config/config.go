package config

import (
	"ai/utils"
	"fmt"
	"os"
	"path"
)

var home, _ = os.UserHomeDir()

var (
	NAME           = "ai"
	VERSION        = "1.0.0"
	CONFIG_DIR     = path.Join(home, ".config", "aicli")
	CONFIG_FILE    = path.Join(CONFIG_DIR, "config.ini")
	LOG_FILE       = path.Join(CONFIG_DIR, "log")
	CONFIG_EXEMPLE = `# Configuration file for aicli`
)

func InitConfig() error {
	tmpName, err := os.MkdirTemp("", "aicli")
	if err != nil {
		return err
	}

	os.Setenv("TMP_DIR", tmpName)

	if !utils.FileExist(CONFIG_DIR) {
		if err := os.MkdirAll(CONFIG_DIR, 0755); err != nil {
			return err
		}
		fmt.Printf("Config dir created at %s\n", CONFIG_DIR)
	}

	if !utils.FileExist(CONFIG_FILE) {
		if err := os.WriteFile(CONFIG_FILE, []byte(CONFIG_EXEMPLE), 0644); err != nil {
			return err
		}
		fmt.Printf("Config file created at %s\n", CONFIG_FILE)
	}

	if !utils.FileExist(LOG_FILE) {
		if _, err := os.Create(LOG_FILE); err != nil {
			return err
		}
		fmt.Printf("Log file created at %s\n", LOG_FILE)
	}

	os.Setenv("LOG_FILE", LOG_FILE)

	return nil
}
