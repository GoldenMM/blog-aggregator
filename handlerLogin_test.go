package main

import (
	"testing"

	"github.com/GoldenMM/blog-aggregator/internal/config"
)

func TestHandlerLogin(t *testing.T) {
	tests := []struct {
		name    string
		cmd     command
		wantErr bool
	}{
		{"ValidLogin1", command{args: []string{"testuser1"}}, false},
		{"NoUsername", command{args: []string{}}, true},
		{"MultipleArgs", command{args: []string{"user1", "user2"}}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &state{cfg: &config.Config{}}

			err := handlerLogin(s, tt.cmd)
			if (err != nil) != tt.wantErr {
				t.Errorf("FAIL [%s]: Expected error %v, got %v", tt.name, tt.wantErr, err != nil)
			}

			// Verify the user was set correctly
			result, err := config.Read()
			if err != nil {
				t.Errorf("FAIL [%s]: Unable to read config file: %v", tt.name, err)
			}

			if !tt.wantErr {
				if result.CurrentUserName != tt.cmd.args[0] {
					t.Errorf("FAIL [%s]: Expected user: [%v], got: [%v]", tt.name, tt.cmd.args[0], s.cfg.CurrentUserName)
				}
			}

		})
	}
}
