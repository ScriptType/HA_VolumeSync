package util

import (
	"bytes"
	"os/exec"
	"strings"
)

func GetVolume() (string, error) {
	cmd := exec.Command("osascript", "-e", "output volume of (get volume settings)")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out.String())), nil
}

func GetMute() (bool, error) {
	output, err := exec.Command("osascript", "-e", "output muted of (get volume settings)").Output()
	if err != nil {
		return false, err
	}

	muteStatus := strings.TrimSpace(string(output))
	return muteStatus == "true", nil
}
