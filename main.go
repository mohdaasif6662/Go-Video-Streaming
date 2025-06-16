package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"github.com/gin-gonic/gin"
)

func main() {
	hlsDir := "./videos"
	if _, err := os.Stat(hlsDir); os.IsNotExist(err) {
		log.Fatalf("HLS directory %s does not exist", hlsDir)
	}

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/" , func (c *gin.Context)  {
		c.HTML(http.StatusOK,"index.tmpl",gin.H{})
	})
	router.GET("/hls/*filepath", func(c *gin.Context) {
		filepathParam := c.Param("filepath")
		// Remove leading slash from the filepath parameter
		requestedFile := filepathParam[1:]
		fullPath := filepath.Join(hlsDir, requestedFile)

		// Security checks
		if filepath.IsAbs(requestedFile) || containsDotDot(requestedFile) {
			c.HTML(http.StatusForbidden,"403.tmpl",gin.H{})
			return
		}

		// Check if the file exists and is not a directory
		fileInfo, err := os.Stat(fullPath)
		if os.IsNotExist(err) || err != nil || fileInfo.IsDir() {
			c.HTML(http.StatusForbidden,"403.tmpl",gin.H{})
			return
		}

		// Verify the file is within the hlsDir
		relPath, err := filepath.Rel(hlsDir, fullPath)
		if err != nil || relPath == ".." || strings.HasPrefix(relPath, "../") {
			c.HTML(http.StatusForbidden,"403.tmpl",gin.H{})
			return
		}

		c.File(fullPath)
	})

	log.Println("Starting server on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// Helper function to check for path traversal attempts
func containsDotDot(v string) bool {
	if !strings.Contains(v, "..") {
		return false
	}
	for _, ent := range strings.FieldsFunc(v, isSlash) {
		if ent == ".." {
			return true
		}
	}
	return false
}

func isSlash(r rune) bool {
	return r == '/' || r == '\\'
}