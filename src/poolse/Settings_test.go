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
		t.Errorf("--- SETTINGS did not match the clone - config file improperly parsed")
	}
}

// TestSettingsRefresh
func TestSettingsRefresh(t *testing.T) {
	SETTINGS = Settings{
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
	SETTINGS.populateTargets()
	if STATUS.Targets[0].DownCountThreshold != 1 ||
		STATUS.Targets[0].UpCountThreshold != 1 {
		t.Errorf("--- Up/Down count threstholds should be set 1 if configured to be less than 1.")
	}
	if STATUS.Targets[0].Endpoint != SETTINGS.Targets[0].Endpoint {
		t.Errorf("SETTINGS and STATUS first target endpoint do not match.")
	}
}

func TestStoppingMonitorProcesses(t *testing.T) {

}

func TestCheckingStartupState(t *testing.T) {
	s := Settings{
		State: State{
			StartupState: true,
		},
		Targets: []Target{
			Target{
				PollingInterval: 1,
				OK:              true,
			},
		},
	}
	STATUS.Targets = []Target{
		Target{
			OK: true,
		},
	}

	s.checkStartupState()
	WG.Wait()

	if !STATUS.State.OK {
		t.Errorf("--- Top level state should be OK (true)")
	}
}
