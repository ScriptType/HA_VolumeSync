package util

import (
	"github.com/ScriptType/HA_VolumeSync/homeassistant"
	"github.com/manifoldco/promptui"
)

// ShowMainMenu displays the main menu and returns the selected option
func ShowMainMenu() (string, error) {
	prompt := promptui.Select{
		Label: "Select Action",
		Items: []string{"Start Program", "Update Configuration", "Exit"},
	}

	_, result, err := prompt.Run()
	return result, err
}

// ShowUpdateConfigMenu displays the update configuration menu and updates the selected field
func ShowUpdateConfigMenu(config *homeassistant.Config) error {
	prompt := promptui.Select{
		Label: "Select Configuration to Update",
		Items: []string{"BaseURL", "AuthToken", "EntityID", "Back to Main Menu"},
	}

	_, result, err := prompt.Run()
	if err != nil {
		return err
	}

	if result == "Back to Main Menu" {
		return nil
	}

	return homeassistant.UpdateConfig(config, result)
}
