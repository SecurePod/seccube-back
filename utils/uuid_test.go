package utils

import "testing"

func TestGenerateUUID(t *testing.T) {
	uuid := GenerateUUID()

	if len(uuid) != 32 {
		t.Errorf("Expected UUID to be 32 characters long, got %d", len(uuid))
	}
}
