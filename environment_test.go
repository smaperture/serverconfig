package serverconfig

import (
	"testing"
)

func TestEnvironment(t *testing.T) {
	err := Environment()
	if err != nil {
		t.Errorf("should be nil, got %v", err)
	}
}

func TestLoadConfig(t *testing.T) {
	tests := []struct {
		id       int
		file     string
		required bool
		expected bool
	}{
		{1, "test.env", true, false},
		{2, "missing.env", false, false},
		{3, "", true, true},
		{4, "missing.env", true, true},
	}

	for _, test := range tests {
		err := loadConfig(test.file, test.required)

		if test.expected && err == nil {
			t.Errorf("%d: should have error, got nil", test.id)
		} else if !test.expected && err != nil {
			t.Errorf("%d: should be nil, got %v", test.id, err)
		}
	}
}
