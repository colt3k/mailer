package update

import (
	"testing"
)

func TestCheckUpdateDoesNotPanic(t *testing.T) {
	// Smoke test: ensure CheckUpdate does not panic with a valid app name.
	// It will be a no-op because the artifactory endpoint is unreachable.
	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Fatalf("CheckUpdate panicked: %v", r)
			}
		}()
		CheckUpdate("mailer")
	}()
}
