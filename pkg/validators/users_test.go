package validators_test

import (
	"simple-server/pkg/validators"
	"testing"
)

func TestValidateLogin(t *testing.T) {
	tests := []struct {
		login   string
		wantErr bool
		errMsg  string
	}{
		{"", true, "login cannot be empty"},
		{"a", true, "login must be at least 3 characters long"},
		{"ab", true, "login must be at least 3 characters long"},
		{"abc", false, ""},
		{"valid_login_123", false, ""},
		{"this_login_is_way_too_long_for_the_system", true, "login must be no longer than 20 characters"},
		{"invalid!login", true, "login can only contain letters, digits, and underscores"},
		{"anotherValidLogin", false, ""},
	}

	for _, tt := range tests {
		err := validators.ValidateLogin(tt.login)
		if (err != nil) != tt.wantErr {
			t.Errorf("ValidateLogin(%q) error = %v, wantErr %v", tt.login, err, tt.wantErr)
		}
		if err != nil && err.Error() != tt.errMsg {
			t.Errorf("ValidateLogin(%q) error = %v, want %v", tt.login, err.Error(), tt.errMsg)
		}
	}
}

func TestValidatePassword(t *testing.T) {
	tests := []struct {
		password string
		wantErr  bool
		errMsg   string
	}{
		{"", true, "password must be at least 8 characters long"},
		{"short", true, "password must be at least 8 characters long"},
		{"alllowercase", true, "password must contain at least one uppercase letter"},
		{"ALLUPPERCASE", true, "password must contain at least one lowercase letter"},
		{"NoDigitsHere!", true, "password must contain at least one digit"},
		{"NoSpecial123", true, "password must contain at least one special character"},
		{"Valid1Password!", false, ""},
		{"Another@Valid1", false, ""},
	}

	for _, tt := range tests {
		err := validators.ValidatePassword(tt.password)
		if (err != nil) != tt.wantErr {
			t.Errorf("ValidatePassword(%q) error = %v, wantErr %v", tt.password, err, tt.wantErr)
		}
		if err != nil && err.Error() != tt.errMsg {
			t.Errorf("ValidatePassword(%q) error = %v, want %v", tt.password, err.Error(), tt.errMsg)
		}
	}
}
