# YouTube Downloader
This is a command-line tool for downloading audio or video files from YouTube.

## Getting Started
These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

## Prerequisites
Make sure you have Go installed on your machine. You can download the latest version of Go from the [official website](https://golang.org/dl/).

You will also need to install the `yt-dlp` command-line utility. You can find instructions for installing `yt-dlp` [here](https://github.com/trizen/yt-dlp).

## Installing
Clone the repository to your local machine:
```
git clone https://github.com/your-username/youtube-downloader.git
```

Change into the project directory:
```
cd youtube-downloader
```

## Usage
To download an audio file from YouTube, run the following command:
```
go run main.go <URL> --audio
```

To download a video file from YouTube, run the following command:
```
go run main.go <URL> --video
```

The downloaded file will be saved to the `C:\Users\%USERNAME%\Downloads` directory with the title of the YouTube video as the file name.

## Built With
- [Go](https://golang.org/) - The programming language used
- [yt-dlp](https://github.com/trizen/yt-dlp) - The command-line utility for downloading YouTube videos


## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
