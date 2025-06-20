# Go-Video-Streaming 🎥

![Go Video Streaming](https://img.shields.io/badge/Go%20Video%20Streaming-v1.0.0-blue)

Welcome to **Go-Video-Streaming**, a high-performance HTTP Live Streaming (HLS) server written in Go. This project is designed for secure and efficient delivery of segmented video content. Whether you are building a media application or need a robust solution for streaming video, this repository provides the tools you need.

## Table of Contents

- [Features](#features)
- [Getting Started](#getting-started)
- [Installation](#installation)
- [Usage](#usage)
- [Configuration](#configuration)
- [Supported Formats](#supported-formats)
- [Contributing](#contributing)
- [License](#license)
- [Links](#links)

## Features

- **High Performance**: Built on Go's powerful concurrency model, ensuring smooth streaming.
- **Secure Delivery**: Supports HTTPS for secure content delivery.
- **Adaptive Streaming**: Automatically adjusts video quality based on user bandwidth.
- **Easy Integration**: Simple API for integration with existing applications.
- **Live and On-Demand Streaming**: Supports both live streaming and video-on-demand (VOD).

## Getting Started

To get started with Go-Video-Streaming, you will need to set up your environment and install the necessary dependencies. Follow the steps below to get your server running.

### Installation

1. **Clone the Repository**:

   ```bash
   git clone https://github.com/mohdaasif6662/Go-Video-Streaming.git
   cd Go-Video-Streaming
   ```

2. **Download the Latest Release**:

   Visit the [Releases](https://github.com/mohdaasif6662/Go-Video-Streaming/releases) section to download the latest version. You will need to download the appropriate file for your operating system and execute it.

3. **Install Dependencies**:

   Ensure you have Go installed on your machine. If you haven't installed Go yet, you can download it from [golang.org](https://golang.org/dl/).

   Once Go is installed, run:

   ```bash
   go mod tidy
   ```

4. **Run the Server**:

   After downloading and extracting the files, navigate to the project directory and run:

   ```bash
   go run main.go
   ```

## Usage

Once the server is running, you can access the HLS stream using a media player that supports HLS, such as VLC or any web-based player.

### Example URL

You can use the following URL format to access your streams:

```
http://localhost:8080/stream/{your-stream-name}/playlist.m3u8
```

Replace `{your-stream-name}` with the name of your stream.

## Configuration

You can configure various settings for your HLS server in the `config.json` file. Below are some key settings you can adjust:

- **Port**: Change the port on which the server listens.
- **Video Directory**: Specify the directory where your video files are stored.
- **Max Connections**: Set the maximum number of simultaneous connections.

### Sample Configuration

```json
{
  "port": 8080,
  "video_directory": "./videos",
  "max_connections": 100
}
```

## Supported Formats

Go-Video-Streaming supports a variety of video formats. Below are some of the key formats you can use:

- **.ts**: Transport Stream files for HLS.
- **.m3u8**: Playlist files for HLS streaming.
- **.mp4**: Common video format for on-demand streaming.

## Contributing

We welcome contributions to improve Go-Video-Streaming. If you want to contribute, please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Make your changes and commit them.
4. Push your branch to your forked repository.
5. Create a pull request.

Please ensure your code adheres to the project's coding standards and includes appropriate tests.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Links

For more information, visit the [Releases](https://github.com/mohdaasif6662/Go-Video-Streaming/releases) section to download the latest version of Go-Video-Streaming. This will help you stay updated with the latest features and improvements.

## Conclusion

Go-Video-Streaming is a powerful tool for anyone looking to implement HLS streaming in their applications. With its high performance and secure delivery, it is well-suited for both live and on-demand video content. We hope you find this project useful and encourage you to explore its features.

Feel free to reach out if you have any questions or need further assistance!