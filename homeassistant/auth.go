package homeassistant

import (
	"github.com/gorilla/websocket"
)

func SendAuth(conn *websocket.Conn, authToken string) error {
	message := map[string]interface{}{
		"type":         "auth",
		"access_token": authToken,
	}
	return conn.WriteJSON(message)
}
