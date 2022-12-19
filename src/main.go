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
	M4AFormat  = "bestaudio[ext=m4a]"
	MP3Format  = "best[ext=mp3]"
	BestFormat = "best"
	MP4Format  = "bestvideo[ext=mp4]+bestaudio[ext=m4a]"
	MP4Format2 = "best[ext=mp4]"
	FileFormat = "C:\\Users\\%USERNAME%\\Downloads\\%(title)s.%(ext)s"
)

func executeCommand(url, flag string) error {
	var cmd *exec.Cmd
	switch flag {
	case AudioFlag:
		cmd = exec.Command("yt-dlp", url, "-f", M4AFormat+"/"+MP3Format+"/"+BestFormat, "-o", FileFormat)
	case VideoFlag:
		cmd = exec.Command("yt-dlp", url, "-f", MP4Format+"/"+MP4Format2+"/"+BestFormat, "-o", FileFormat)
	default:
		return fmt.Errorf("unexpected flag: %s", flag)
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("Error al obtener la tubería de salida: %s\n", err)
	}

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("Error al iniciar el comando: %s\n", err)
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
