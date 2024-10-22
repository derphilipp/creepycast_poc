package main

import (
    "fmt"
    "log"
    "os"
    "time"
)

var accessLog *log.Logger

func init() {
    // Open the log file in append mode, create it if it doesn't exist
    file, err := os.OpenFile("access.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatalf("Failed to open log file: %v", err)
    }

    // Create a new logger
    accessLog = log.New(file, "", log.LstdFlags)
}

// Function to log access details
func logAccess(userAgent, ip string) {
    currentTime := time.Now().Format(time.RFC1123)
    logEntry := fmt.Sprintf("Time: %s, IP: %s, User Agent: %s", currentTime, ip, userAgent)
    accessLog.Println(logEntry)
}

