package main

import (
	"fmt"
	"strings"
	"time"
)

// We determine the OS from the User Agent;
// This is a very simple implementation and should not be used in production
// Every OS is mapped to a specific MP3 file
func determineOS(userAgent string) string {
	if strings.Contains(userAgent, "Windows") {
		return "variations/os/windows.mp3"
	} else if strings.Contains(userAgent, "Macintosh") ||
		strings.Contains(userAgent, "darwin") ||
		strings.Contains(userAgent, "Darwin") {
		return "variations/os/macos.mp3"
	} else if strings.Contains(userAgent, "Linux") {
		return "variations/os/linux.mp3"
	} else if strings.Contains(userAgent, "iPhone") ||
		strings.Contains(userAgent, "iPad") ||
		strings.Contains(userAgent, "iOS") {
		return "variations/os/iphone.mp3"
	} else if strings.Contains(userAgent, "Android") {
		return "variations/os/android.mp3"
	}
	return "variations/os/unknown.mp3"
}

// We determine the podcast player (or browser) from the User Agent;
// This is a very simple implementation and should not be used in production
// Also: Often we can deduce the OS from the player;
// If we would do this way better, we should have a look how podlove publisher does it!
//
// Every user agent has an mp3 file
func determineCatcher(userAgent string) string {
	fmt.Printf("User Agent: %s\n", userAgent)
	if strings.Contains(userAgent, "iTunes") || strings.Contains(userAgent, "AppleCoreMedia") {
		return "variations/catcher/apple_podcasts.mp3"
	} else if strings.Contains(userAgent, "Spotify") {
		return "variations/catcher/spotify.mp3"
	} else if strings.Contains(userAgent, "Google") {
		return "variations/catcher/google_podcasts.mp3"
	} else if strings.Contains(userAgent, "Pocket Casts") {
		return "variations/catcher/pocket_casts.mp3"
	} else if strings.Contains(userAgent, "Podcat") {
		return "variations/catcher/podcat.mp3"
	} else if strings.Contains(userAgent, "Podcast Addict") {
		return "variations/catcher/podcast_addict.mp3"
	} else if strings.Contains(userAgent, "Stitcher") {
		return "variations/catcher/stitcher.mp3"
	} else if strings.Contains(userAgent, "TuneIn") {
		return "variations/catcher/tunein.mp3"
	} else if strings.Contains(userAgent, "Chaos box") {
		return "variations/catcher/chaos_box.mp3"
	} else if strings.Contains(userAgent, "Edg") {
		return "variations/catcher/browser_edge.mp3"
	} else if strings.Contains(userAgent, "Chrome") {
		return "variations/catcher/browser_chrome.mp3"
	} else if strings.Contains(userAgent, "Firefox") {
		return "variations/catcher/browser_firefox.mp3"
	} else if strings.Contains(userAgent, "Safari") {
		return "variations/catcher/browser_safari.mp3"
	} else if strings.Contains(userAgent, "wget") || strings.Contains(userAgent, "curl") {
		return "variations/catcher/wget_curl.mp3"
	}
	return "variations/catcher/unknown.mp3"
}

// Same for the hours of the day; we have 00.mp3 to 23.mp3, which we can use
// Be aware: This is running on the server, so we use the localtime of the server, not the client
func determineTime() string {
	hour := time.Now().In(time.Local).Hour()
	fmt.Printf("Time now: %v\n", hour)
	return fmt.Sprintf("variations/time/%02d.mp3", hour)
}
