package song

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func Convert(inputFile string, outputFormat string, freq int, bitrate int, channels int) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(dir)
	command := exec.Command(
		"ffmpeg", "-i", inputFile, "-vn", "-ar", strconv.Itoa(freq), "-ac",
		strconv.Itoa(channels), "-b:a", strconv.Itoa(bitrate)+"k",
		"../converted/"+strings.Split(inputFile, ".")[0]+"."+outputFormat)
	command.Dir = dir + "/song_data/raw/"
	command.Stderr = os.Stderr
	command.Stdout = os.Stdout

	cmdErr := command.Run()
	if cmdErr != nil {
		log.Fatal(err)
	}
}
