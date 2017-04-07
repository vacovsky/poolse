package main

import (
	"os"
	"testing"
)

// TestLoadingSettings
func TestLoadingSettings(t *testing.T) {
	os.Args = []string{
		"testing",
		"config_test.json",
	}
	SETTINGS.parseSettingsFile()

	settingsShouldBeThis := Settings{
		Targets: []Target{
			Target{
				Endpoint:           "beer://thisisfortestingonly.whatever",
				PollingInterval:    -1,
				ExpectedStatusCode: -1,
				DownCountThreshold: -1,
				UpCountThreshold:   -1,
			},
		},
	}
	settingsShouldBeThis.Service.Debug = true
	settingsShouldBeThis.Service.HTTPPort = "1234567"
	settingsShouldBeThis.Service.StateFileName = "indialpaleale.dat"

	pass := (settingsShouldBeThis.State == SETTINGS.State &&
		settingsShouldBeThis.Targets[0].Endpoint == SETTINGS.Targets[0].Endpoint &&
		settingsShouldBeThis.Targets[0].PollingInterval == SETTINGS.Targets[0].PollingInterval &&
		settingsShouldBeThis.Targets[0].ExpectedStatusCode == SETTINGS.Targets[0].ExpectedStatusCode &&
		settingsShouldBeThis.Targets[0].UpCountThreshold == SETTINGS.Targets[0].UpCountThreshold)

	if !pass {
		t.Errorf("SETTINGS did not match the clone - config file improperly parsed")
	}
}

// TestSettingsRefresh
func TestSettingsRefresh(t *testing.T) {

}
