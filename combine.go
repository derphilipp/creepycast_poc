package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/exec"
	"time"
)

// Here we use ffmpeg to combine the mp3 files
// ffmpeg has a concat filter that can be used to concatenate multiple files
// We comebine the 'time of the day' mp3
// with the 'generic' mp3 which is the same for all users at all times
// then the different catchers
// and finally different os

// Imagine what "fun" we could have it we would geolocate the users and
// start to play a specific file for their city or country...

func loadAudioFile(w http.ResponseWriter, r *http.Request) {
	// Collect information about the request
	userAgent := r.Header.Get("User-Agent")
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	logAccess(userAgent, ip)

	// Determine the dynamic parts
	osFile := determineOS(userAgent)
	catcherFile := determineCatcher(userAgent)
	timeFile := determineTime()
	genericFile := "variations/generic/generic.mp3"

	// Print the collected information
	fmt.Printf("User Agent: %s\n", userAgent)
	fmt.Printf("IP Address: %s\n", ip)
	fmt.Printf("OS File: %s\n", osFile)
	fmt.Printf("Catcher File: %s\n", catcherFile)
	fmt.Printf("Time File: %s\n", timeFile)

	// Combine the MP3 files into a single temporary file
	tempFile := fmt.Sprintf("temp_%d.mp3", time.Now().UnixNano())
	cmd := exec.Command("ffmpeg", "-y", "-i", "concat:"+timeFile+"|"+genericFile+"|"+catcherFile+"|"+osFile, "-c", "copy", tempFile)
	if err := cmd.Run(); err != nil {
		http.Error(w, "Error combining files.", http.StatusInternalServerError)
		return
	}
	defer os.Remove(tempFile)

	// Serve the combined MP3 file
	http.ServeFile(w, r, tempFile)
}
