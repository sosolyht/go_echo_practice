package handler

import (
	"encoding/json"
	"log"
	"os/exec"
)

type good struct{}

func FFprobe(url string) map[string]interface{} {
	result := exec.Command("ffprobe", "-show_streams", "-show_format", "-print_format", "json", url)
	result1, err := result.Output()
	if err != nil {
		log.Fatal(err)
	}
	result2 := string(result1)
	var a map[string]interface{}
	json.Unmarshal([]byte(result2), &a)

	return a
}
