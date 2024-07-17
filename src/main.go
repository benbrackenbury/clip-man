package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/atotto/clipboard"
)

func main() {
	var lastContent string

	// Create or open log file
	file, err := os.OpenFile("clipboard.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	logger := log.New(file, "", log.LstdFlags)

	// Setup signal catching
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Monitor clipboard
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	fmt.Println("Monitoring clipboard. Press Ctrl+C to stop.")
	for {
		select {
		case <-ticker.C:
			content, err := clipboard.ReadAll()
			if err != nil {
				log.Printf("Error reading clipboard: %v", err)
				continue
			}

			if content != lastContent {
				logger.Printf("Clipboard changed: %s", content)
				lastContent = content
			}

		case sig := <-sigs:
			fmt.Printf("\nReceived signal %s. Exiting...\n", sig)
			return
		}
	}
}

