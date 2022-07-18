package handler

import (
	"encoding/json"
	"os/exec"
)

type result struct {
	Streams string `json:"streams"`
}

func FFprobe(url string) *result {
	var a result
	cmd := exec.Command("ffprobe", "-show_streams", "-show_format", "-print_format", "json", url)
	cmdOutput, err := cmd.Output()
	if err != nil {
		log.Error(err)
	}
	cmdStr := string(cmdOutput)
	json.Unmarshal([]byte(cmdStr), &a)

	return &a
}
