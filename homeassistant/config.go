package homeassistant

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// Config holds the configuration values
type Config struct {
	BaseURL            string `json:"base_url"`
	HomeAssistantWSURL string `json:"-"`
	HomeAssistantURL   string `json:"-"`
	AuthToken          string `json:"auth_token"`
	EntityID           string `json:"entity_id"`
}

// LoadConfig reads configuration from a file or creates one if it doesn't exist
func LoadConfig(filename string) (*Config, error) {
	var config Config
	file, err := os.Open(filename)
	if os.IsNotExist(err) {
		fmt.Println("Configuration file not found, creating a new one.")
		config, err = createConfig(filename)
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	} else {
		defer file.Close()
		bytes, err := io.ReadAll(file)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(bytes, &config)
		if err != nil {
			return nil, err
		}
	}

	// Set derived values
	config.HomeAssistantWSURL = "ws://" + config.BaseURL + "/api/websocket"
	config.HomeAssistantURL = "http://" + config.BaseURL + "/api"

	fmt.Println(config.HomeAssistantWSURL)

	return &config, nil
}

// SaveConfig writes the configuration to a file
func SaveConfig(filename string, config *Config) error {
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0o644)
}

// createConfig creates a new configuration file with user input
func createConfig(filename string) (Config, error) {
	var config Config
	fmt.Print("Enter Base URL: ")
	fmt.Scanln(&config.BaseURL)
	fmt.Print("Enter Auth Token: ")
	fmt.Scanln(&config.AuthToken)
	fmt.Print("Enter Entity ID: ")
	fmt.Scanln(&config.EntityID)

	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return config, err
	}

	err = os.WriteFile(filename, data, 0o644)
	if err != nil {
		return config, err
	}

	return config, nil
}

// UpdateConfig updates specific fields of the configuration
func UpdateConfig(config *Config, field string) error {
	switch field {
	case "BaseURL":
		fmt.Print("Enter new Base URL: ")
		fmt.Scanln(&config.BaseURL)
	case "AuthToken":
		fmt.Print("Enter new Auth Token: ")
		fmt.Scanln(&config.AuthToken)
	case "EntityID":
		fmt.Print("Enter new Entity ID: ")
		fmt.Scanln(&config.EntityID)
	default:
		return fmt.Errorf("unknown configuration field")
	}

	config.HomeAssistantWSURL = config.BaseURL + "/api/websocket"
	config.HomeAssistantURL = config.BaseURL + "/api"

	return nil
}
