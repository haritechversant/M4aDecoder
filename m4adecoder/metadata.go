package m4adecoder

import (
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

// GetAudioMetadata extracts the sample rate and duration of an audio file.
func GetAudioMetadata(filePath string) (sampleRate int, duration float64, err error) {
	cmd := exec.Command("ffprobe", "-v", "error", "-select_streams", "a:0", "-show_entries", "stream=sample_rate,duration", "-of", "default=noprint_wrappers=1:nokey=1", filePath)
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		return 0, 0, fmt.Errorf("error running ffprobe: %v", err)
	}

	// Parse the output
	data := strings.TrimSpace(out.String())
	parts := strings.Split(data, "\n")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("unexpected output format: %s", data)
	}

	// Convert the values to appropriate types
	sampleRate, err = strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, fmt.Errorf("error converting sample rate: %v", err)
	}

	duration, err = strconv.ParseFloat(parts[1], 64)
	if err != nil {
		return 0, 0, fmt.Errorf("error converting duration: %v", err)
	}

	return sampleRate, duration, nil
}
