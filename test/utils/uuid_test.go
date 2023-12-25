package test

import (
	. "docker-api/utils"
	"strings"
	"testing"
)

func TestGenerateUUID(t *testing.T) {
	tests := []struct {
		name        string
		expectedLen int
	}{
		{"ValidUUIDLength", 32},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uuid := GenerateUUID()
			if len(uuid) != tt.expectedLen {
				t.Errorf("GenerateUUID() length = %v, want %v", len(uuid), tt.expectedLen)
			}
			if strings.Contains(uuid, "-") {
				t.Errorf("GenerateUUID() should not contain hyphens, got %v", uuid)
			}
		})
	}
}
