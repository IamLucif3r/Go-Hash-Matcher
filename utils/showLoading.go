package utils

import (
	"fmt"
	"time"
)

func ShowLoadingSpinner() {
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	spinnerFrames := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}

	i := 0
	for {
		select {
		case <-ticker.C:
			fmt.Printf("\rChecking passwords %s", spinnerFrames[i])
			i = (i + 1) % len(spinnerFrames)
		}
	}
}
