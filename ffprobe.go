package main

import (
	"log"
	"os/exec"
	"strconv"
	"strings"
)

func execFfprobe(filePath string) string {
	cmd := exec.Command("ffprobe", "-v", "error", "-show_entries", "format=duration", "-of", "default=noprint_wrappers=1:nokey=1", filePath)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Failed to run ffprobe: %v", err)
	}

	return strings.TrimSpace(string(out))
}

func GetDuration(filePath string) float64 {
	d := execFfprobe(filePath)

	duration, err := strconv.ParseFloat(d, 64)
	if err != nil {
		log.Fatalf("Failed to get duration: %v", err)
	}

	return duration
}
