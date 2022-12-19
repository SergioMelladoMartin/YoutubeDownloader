package main

import (
	"fmt"
	"os"
	"os/exec"
	"io"
)

const (
	AudioFlag  = "--audio"
	VideoFlag  = "--video"
	MP3Format  = "bestaudio[ext=m4a]/best[ext=mp3]/best"
	MP4Format  = "bestvideo[ext=m4a]/best[ext=mp4]/best"
	File = "C:\\Users\\%USERNAME%\\Downloads\\%(title)s.%(ext)s"
)

func executeCommand(url, flag string) error {
	var cmd *exec.Cmd
	switch flag {
	  case AudioFlag:
      cmd = exec.Command("yt-dlp", url, "-f", MP3Format, "-o", File)
    case VideoFlag:
		  cmd = exec.Command("yt-dlp", url, "-f", MP4Format, "-o", File)
	  default:
		  return fmt.Errorf("unexpected flag: %s", flag)
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("Error while creating StdoutPipe for Cmd: %s\n", err)
	}

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("Error while starting Cmd: %s\n", err)
	}

	buf := make([]byte, 1024)
	for {
		n, err := stdout.Read(buf)
		if n > 0 {
			fmt.Print(string(buf[:n]))
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("Error: %s\n", err)
		}
	}

	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("Error: %s\n", err)
	}

	return nil
}

func main() {
	if len(os.Args) < 3 || (os.Args[2] != AudioFlag && os.Args[2] != VideoFlag) {
		fmt.Println("Usage: main <url> <--audio/--video>")
		return
	}

	url := os.Args[1]
	flag := os.Args[2]

	if err := executeCommand(url, flag); err != nil {
		fmt.Println(err)
		return
	}
}
