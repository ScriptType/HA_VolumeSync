package homeassistant

import (
	"bytes"
	"fmt"
	"net/http"
)

func ChangeVolume(entityID string, volumeLevel float64, homeAssistantURL, authToken string) error {
	body := fmt.Sprintf(`{"entity_id": "%s", "volume_level": %f}`, entityID, volumeLevel)
	req, err := http.NewRequest("POST", homeAssistantURL+"/services/media_player/volume_set", bytes.NewBufferString(body))
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+authToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received non-ok HTTP status: %s", resp.Status)
	}

	fmt.Println("Volume changed successfully")
	return nil
}
