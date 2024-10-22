package main

import (
	"fmt"
	"net/http"
	"time"
)

// Function to print the current time and timezone
func printCurrentTimeAndTimezone() {
    currentTime := time.Now().In(time.Local)
    fmt.Printf("Current Time: %s\n", currentTime.Format(time.RFC1123))
    fmt.Printf("Timezone: %s\n", currentTime.Location())
}


// This is a HORRIBLE, but working EXTREMLY simpel RSS feed
func generateRSSFeed() string {
	return `<?xml version="1.0" encoding="UTF-8" ?>
<rss version="2.0">
    <channel>
        <title>Creepycast - the generated Podcast</title>
        <link>https://experiment.grumhold.de/</link>
        <description>This is a sample podcast feed</description>
        <item>
            <title>Episode 1</title>
            <enclosure url="https://experiment.grumhold.de/episode.mp3" type="audio/mpeg"/>
            <guid>1</guid>
            <pubDate>Mon, 01 Jan 2023 00:00:00 +0000</pubDate>
        </item>
    </channel>
</rss>`
}

// We serve an rss feed (see above) and an audio file
// To see the combinations look at 'variations.go' and 'combine.go'
func main() {
        printCurrentTimeAndTimezone()

	http.HandleFunc("/rss", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/rss+xml")
		fmt.Fprint(w, generateRSSFeed())
	})

	http.HandleFunc("/episode.mp3", loadAudioFile)

	fmt.Println("Starting server at :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
