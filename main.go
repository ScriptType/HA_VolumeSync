package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/ScriptType/HA_VolumeSync/homeassistant"
	"github.com/ScriptType/HA_VolumeSync/keyboard"
	"github.com/ScriptType/HA_VolumeSync/keylogger"
	"github.com/ScriptType/HA_VolumeSync/util"
	"github.com/gorilla/websocket"
)

var (
	syncVolume = false
	lastVolume = "0.5"
	config     *homeassistant.Config
)

func main() {
	var err error
	config, err = homeassistant.LoadConfig("config.json")
	if err != nil {
		log.Fatal(err)
	}

	for {
		choice, err := util.ShowMainMenu()
		if err != nil {
			log.Fatal(err)
		}

		switch choice {
		case "Start Program":
			startProgram()
		case "Update Configuration":
			if err := util.ShowUpdateConfigMenu(config); err != nil {
				log.Fatal(err)
			}
			if err := homeassistant.SaveConfig("config.json", config); err != nil {
				log.Fatal(err)
			}
		case "Exit":
			fmt.Println("Exiting...")
			return
		}
	}
}

func startProgram() {
	kl, conn, err := initialize()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	go kl.Listen(handleKeyEvent)

	for {
		if err := readWebSocketMessages(conn); err != nil {
			log.Println("Error reading WebSocket message:", err)
			break
		}
	}
}

func initialize() (*keylogger.KeyLogger, *websocket.Conn, error) {
	kl, err := keylogger.New()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to initialize keylogger: %w", err)
	}

	conn, err := connectToWebSocket()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to WebSocket: %w", err)
	}

	if err := authenticateWebSocket(conn); err != nil {
		return nil, nil, fmt.Errorf("failed to authenticate WebSocket: %w", err)
	}

	return kl, conn, nil
}

func connectToWebSocket() (*websocket.Conn, error) {
	headers := http.Header{}
	headers.Add("Authorization", "Bearer "+config.AuthToken)

	u, err := url.Parse(config.HomeAssistantWSURL)
	if err != nil {
		return nil, fmt.Errorf("error parsing URL: %w", err)
	}

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), headers)
	if err != nil {
		return nil, fmt.Errorf("error connecting to WebSocket: %w", err)
	}

	return conn, nil
}

func authenticateWebSocket(conn *websocket.Conn) error {
	if err := homeassistant.SendAuth(conn, config.AuthToken); err != nil {
		return fmt.Errorf("authentication failed: %w", err)
	}

	if err := homeassistant.SubscribeToStateChanges(conn); err != nil {
		return fmt.Errorf("subscription failed: %w", err)
	}

	return nil
}

func handleKeyEvent(key keyboard.Key, state keyboard.State) {
	if key == keyboard.MediaPrevious && syncVolume && state == keyboard.Up {
		go homeassistant.PreviousTrack(config.EntityID, config.HomeAssistantURL, config.AuthToken)
	}
	if key == keyboard.MediaPlayPause && syncVolume && state == keyboard.Up {
		fmt.Println("MediaPlayPause")
	}
	if key == keyboard.MediaNext && syncVolume && state == keyboard.Up {
		go homeassistant.NextTrack(config.EntityID, config.HomeAssistantURL, config.AuthToken)
	}
	if key == keyboard.MediaMute && syncVolume && state == keyboard.Down {
		handleMute(config.EntityID)
	}
	if (key == keyboard.MediaVolumeDown || key == keyboard.MediaVolumeUp) && syncVolume && state == keyboard.Up {
		changeVolume(config.EntityID)
	}
}

func handleMute(entityID string) {
	muteStat, err := util.GetMute()
	if err != nil {
		fmt.Println("Error getting mute status:", err)
		return
	}
	if !muteStat {
		go homeassistant.ChangeVolume(entityID, 0.0, config.HomeAssistantURL, config.AuthToken)
	} else {
		parsedVolume, err := strconv.ParseFloat(lastVolume, 64)
		if err != nil {
			fmt.Println("Error parsing volume:", err)
		} else {
			go homeassistant.ChangeVolume(entityID, parsedVolume/100, config.HomeAssistantURL, config.AuthToken)
		}
	}
}

func changeVolume(entityID string) {
	volume, err := util.GetVolume()
	if err != nil {
		fmt.Println("Error getting volume:", err)
		return
	}
	if volume != lastVolume {
		parsedVolume, err := strconv.ParseFloat(volume, 64)
		if err != nil {
			fmt.Println("Error parsing volume:", err)
		} else {
			go homeassistant.ChangeVolume(entityID, parsedVolume/100, config.HomeAssistantURL, config.AuthToken)
			fmt.Printf("Volume Changed to: %s\n", volume)
			lastVolume = volume
		}
	}
}

func readWebSocketMessages(conn *websocket.Conn) error {
	_, message, err := conn.ReadMessage()
	if err != nil {
		return err
	}
	homeassistant.CheckPlayerState(message, config.EntityID, &syncVolume)
	return nil
}
