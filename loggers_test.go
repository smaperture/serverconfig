package serverconfig

import (
	"testing"

	"github.com/designsbysm/timber/v2"
	"github.com/spf13/viper"
)

func setupLoggers() {
	timber.Reset()
	viper.Reset()
}

func TestLoggers(t *testing.T) {
	setupLoggers()

	err := Loggers()
	if err != nil {
		t.Errorf("should have nil, got %v", err)
	}
}

func TestLoggerSetup(t *testing.T) {
	type Setting struct {
		key   string
		value string
	}

	tests := []struct {
		id         int
		fn         func() error
		haveError  bool
		haveWriter bool
		settings   []Setting
	}{
		{1, cli, false, true, []Setting{
			{"timber.cli.level", "LevelAll"},
		}},
		{2, cli, false, false, []Setting{
			{"timber.cli.level", ""},
		}},
		{3, email, true, false, []Setting{
			{"timber.email.level", "LevelAll"},
		}},
		{4, email, false, false, []Setting{
			{"timber.email.level", ""},
		}},
		{5, email, false, true, []Setting{
			{"timber.email.level", "LevelAll"},
			{"timber.email.host", "test.com"},
			{"timber.email.port", "443"},
			{"timber.email.subject", "Testing"},
			{"timber.email.from", "user@test.com"},
			{"timber.email.to", "user@test.com"},
		}},
		{6, file, true, false, []Setting{
			{"timber.file.level", "LevelAll"},
		}},
		{7, file, false, false, []Setting{
			{"timber.file.level", ""},
		}},
		{8, file, false, true, []Setting{
			{"timber.file.level", "LevelAll"},
			{"timber.file.path", "logs"},
		}},
	}

	for _, test := range tests {
		setupLoggers()

		for _, setting := range test.settings {
			viper.Set(setting.key, setting.value)
		}

		err := test.fn()
		if test.haveError && err == nil {
			t.Errorf("%d: should have error, got nil", test.id)
		} else if !test.haveError && err != nil {
			t.Errorf("%d: should be nil, got %v", test.id, err)
		}

		writers := len(timber.Writers())
		if test.haveWriter && writers == 0 {
			t.Errorf("%d: should have writers, got empty", test.id)
		} else if !test.haveWriter && writers > 0 {
			t.Errorf("%d: should be empty, got %d", test.id, writers)
		}
	}
}
