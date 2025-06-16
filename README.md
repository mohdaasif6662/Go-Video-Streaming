# Go Video Streaming - Lightweight HLS Streaming Server  

![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go) 
![Gin](https://img.shields.io/badge/Gin-1.9+-000000?logo=go) 
![License](https://img.shields.io/badge/License-MIT-blue)

A high-performance HTTP Live Streaming (HLS) server written in Go, designed for secure and efficient delivery of segmented video content.

## Features

- 🚀 **Fast Media Delivery**: Built with Go and Gin for high-throughput streaming
- 🔒 **Security-First**: Protects against path traversal attacks with strict file access controls
- 📁 **Simple Setup**: Just point to your HLS directory (`m3u8` playlists + `ts` segments)
- 🛡️ **Auto-Error Handling**: Custom 403/404 pages for forbidden/missing content
- 🌐 **Web Interface**: Optional homepage template for status/instructions

## Use Cases

- Private video streaming platforms
- Low-latency live event broadcasting
- Secure internal media distribution
- Development/testing for HLS applications

## Quick Start

```bash
# Clone repository
git clone https://github.com/majdsassi/Go-Video-Streaming.git
cd Go-Video-Streaming

# Create videos directory (or mount your existing HLS content)
mkdir -p videos/stream1

# Add HLS files (e.g., videos/stream1/master.m3u8 + segments)

# Run server (default port :8080)
go run main.go
