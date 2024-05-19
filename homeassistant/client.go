package homeassistant

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

func SubscribeToStateChanges(conn *websocket.Conn) error {
	message := map[string]interface{}{
		"id":         1,
		"type":       "subscribe_events",
		"event_type": "state_changed",
	}
	return conn.WriteJSON(message)
}

func CheckPlayerState(message []byte, entityID string, syncVolume *bool) {
	var msg map[string]interface{}
	if err := json.Unmarshal(message, &msg); err != nil {
		log.Println("Error decoding message:", err)
		return
	}

	if msg["type"] == "event" {
		event := msg["event"].(map[string]interface{})
		data := event["data"].(map[string]interface{})
		entity := data["entity_id"].(string)
		if entity == entityID {
			newState := data["new_state"].(map[string]interface{})
			state := newState["state"].(string)
			if (state == "playing" || state == "paused") && !*syncVolume {
				*syncVolume = true
				fmt.Println("Syncing volume")
			} else if state != "playing" && state != "paused" {
				*syncVolume = false
				fmt.Println("Not syncing volume")
			}
		}
	}
}
