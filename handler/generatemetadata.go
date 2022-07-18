package handler

import (
	"encoding/json"
	"log"
	"os/exec"
)

func FFprobe(url string) map[string]interface{} {
	var result map[string]interface{}

	cmd := exec.Command("ffprobe", "-show_streams", "-show_format", "-print_format", "json", url)
	cmdOutput, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	cmdStr := string(cmdOutput)
	json.Unmarshal([]byte(cmdStr), &result)

	return result
}
