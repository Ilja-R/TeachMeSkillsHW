package configs

import (
	"encoding/json"
	"fmt"
	"github.com/Ilja-R/TeachMeSkillsHW/project-1/internal/models"
	"github.com/joho/godotenv"
	"os"
)

var AppSettings models.Config

func ReadSettings() error {
	if err := godotenv.Load(".env"); err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}

	configFile, err := os.Open("internal/configs/configs.json")
	if err != nil {
		return fmt.Errorf("error while opening config file: %w", err)
	}
	defer configFile.Close()

	if err = json.NewDecoder(configFile).Decode(&AppSettings); err != nil {
		return fmt.Errorf("error while parsing config file: %w", err)
	}

	return nil
}
