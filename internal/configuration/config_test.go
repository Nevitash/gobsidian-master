package configuration
package configuration

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetConfig(t *testing.T) {
	// Test initial state
	assert.Nil(t, GetConfig(), "Initial config should be nil")

	// Test after setting config
	testConfig := &Config{
		ConfigPath: "test.yaml",
		VaultPath:  "test/vault",
	}
	SetConfig(testConfig)
	assert.Equal(t, testConfig, GetConfig(), "GetConfig should return the config that was set")
}

func TestGetIncludeGlob(t *testing.T) {
	tests := []struct {
		name           string
		includePattern []string
		wantErr        bool
	}{
		{
			name:           "empty patterns",
			includePattern: []string{},
			wantErr:        true,
		},
		{
			name:           "single pattern",
			includePattern: []string{"*.md"},
			wantErr:        false,
		},
		{
			name:           "multiple patterns",
			includePattern: []string{"*.md", "*.txt"},
			wantErr:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				IncludePatterns: tt.includePattern,
			}
			glob, err := c.GetIncludeGlob()
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, glob)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, glob)
			}
		})
	}
}

func TestGetExcludeGlob(t *testing.T) {
	tests := []struct {
		name           string
		excludePattern []string
		wantErr        bool
	}{
		{
			name:           "empty patterns",
			excludePattern: []string{},
			wantErr:        true,
		},
		{
			name:           "single pattern",
			excludePattern: []string{".git/*"},
			wantErr:        false,
		},
		{
			name:           "multiple patterns",
			excludePattern: []string{".git/*", "node_modules/*"},
			wantErr:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				ExcludePatterns: tt.excludePattern,
			}
			glob, err := c.GetExcludeGlob()
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, glob)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, glob)
			}
		})
	}
}