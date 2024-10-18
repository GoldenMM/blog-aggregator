package config

import (
	"reflect"
	"testing"
)

func TestRead(t *testing.T) {

	var tests = []struct {
		name    string
		want    Config
		errWant error
	}{
		{"TestRead", Config{}, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Read()
			if err != tt.errWant {
				t.Errorf("FAIL [%s]: Expected error %v, got %v", tt.name, tt.errWant, err)
			}

			if reflect.TypeOf(result) != reflect.TypeOf(tt.want) {
				t.Errorf("FAIL [%s]: Expected %v, got %v", tt.name, tt.want, result)
			}
		})
	}
}
func TestGetConfigFilePath(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"ValidHomeDir", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := getConfigFilePath()
			if (err != nil) != tt.wantErr {
				t.Errorf("FAIL [%s]: Expected error %v, got %v", tt.name, tt.wantErr, err != nil)
			}
		})
	}
}
func TestWrite(t *testing.T) {
	// Restore file to default after the test
	defer write(Config{DbURL: "postgres://example", CurrentUserName: "testuser"})

	tests := []struct {
		name    string
		cfg     Config
		wantErr bool
	}{
		{"ValidConfig", Config{DbURL: "http://localhost:5432", CurrentUserName: "testuser"}, false},
		{"InvalidConfig", Config{DbURL: "", CurrentUserName: ""}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := write(tt.cfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("FAIL [%s]: Expected error %v, got %v", tt.name, tt.wantErr, err != nil)
			}

			if !tt.wantErr {
				// Verify the file was written correctly
				result, err := Read()
				if err != nil {
					t.Errorf("FAIL [%s]: Unable to read config file: %v", tt.name, err)
				}

				if !reflect.DeepEqual(result, tt.cfg) {
					t.Errorf("FAIL [%s]: Expected config %v, got %v", tt.name, tt.cfg, result)
				}
			}
		})
	}
}
func TestSetUser(t *testing.T) {
	// Restore file to default after the test
	defer write(Config{DbURL: "postgres://example", CurrentUserName: "testuser"})

	tests := []struct {
		name    string
		user    string
		wantErr bool
	}{
		{"ValidUser", "newuser", false},
		{"EmptyUser", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// Call SetUser
			err := Config{}.SetUser(tt.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("FAIL [%s]: Expected error %v, got %v", tt.name, tt.wantErr, err != nil)
			}

			// Verify the user was set correctly
			if !tt.wantErr {
				result, err := Read()
				if err != nil {
					t.Errorf("FAIL [%s]: Unable to read config file: %v", tt.name, err)
				}

				if result.CurrentUserName != tt.user {
					t.Errorf("FAIL [%s]: Expected user %v, got %v", tt.name, tt.user, result.CurrentUserName)
				}
			}
		})
	}
}
