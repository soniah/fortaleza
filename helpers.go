package fortaleza

import "log"

// Check provides quick error handling
func Check(err error) {
	if err != nil {
		log.Fatalf("Check: %v", err)
	}
}
