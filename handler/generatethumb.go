package handler

import (
	"fmt"
	"os/exec"
)

func FFmpeg(url string) string {
	data := fmt.Sprintf(`"%s"`, url)
	log.Info("Get url")
	cmd := exec.Command("ffmpeg", "-i", data, "-vcodec", "libwebp", "-preset", "default", "-loop", "0", "-an", "-vsync", "0", "-vf", `"fps=10,scale=854:480"`, "-qscale", "10", "-frames:v", "1", "output.webp")
	log.Info("create webp")
	cmd.Run()
	log.Info("command run")

	fmt.Println(cmd)

	//cmdOutput, err := cmd.Output()
	//fmt.Println("\n", cmd)
	//if err != nil {
	//	log.Error(err)
	//}
	//cmdStr := string(cmdOutput)
	//fmt.Println(cmdStr)
	//json.Unmarshal([]byte(cmdStr), &result)

	return "wait"
}
