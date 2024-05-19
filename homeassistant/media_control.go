package homeassistant

import (
	"bytes"
	"fmt"
	"net/http"
)

func sendMediaControl(entityID, service, homeAssistantURL, authToken string) error {
	body := fmt.Sprintf(`{"entity_id": "%s"}`, entityID)
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/services/media_player/%s", homeAssistantURL, service), bytes.NewBufferString(body))
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

	fmt.Printf("%s command sent successfully\n", service)
	return nil
}

func PreviousTrack(entityID, homeAssistantURL, authToken string) error {
	return sendMediaControl(entityID, "media_previous_track", homeAssistantURL, authToken)
}

func NextTrack(entityID, homeAssistantURL, authToken string) error {
	return sendMediaControl(entityID, "media_next_track", homeAssistantURL, authToken)
}

func Play(entityID, homeAssistantURL, authToken string) error {
	return sendMediaControl(entityID, "media_play", homeAssistantURL, authToken)
}

func Pause(entityID, homeAssistantURL, authToken string) error {
	return sendMediaControl(entityID, "media_pause", homeAssistantURL, authToken)
}
