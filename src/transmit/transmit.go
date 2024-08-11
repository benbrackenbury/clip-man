package transmit

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/atotto/clipboard"
	"github.com/benbrackenbury/clipman/src/store"
)

func Transmit(store store.Store) {
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

            store.SetLatestContent(content)

		case sig := <-sigs:
			fmt.Printf("\nReceived signal %s. Exiting...\n", sig)
			return
		}
	}
}

