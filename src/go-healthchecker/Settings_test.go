package main

/* TestConfigFileLoadsCorrectly
func TestSomeSettingsStuff(t *testing.T) {
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
		settingsShouldBeThis.Targets[1].Endpoint == SETTINGS.Targets[1].Endpoint &&
		settingsShouldBeThis.Targets[1].PollingInterval == SETTINGS.Targets[1].PollingInterval &&
		settingsShouldBeThis.Targets[1].ExpectedStatusCode == SETTINGS.Targets[1].ExpectedStatusCode &&
		settingsShouldBeThis.Targets[1].UpCountThreshold == SETTINGS.Targets[1].UpCountThreshold)

	if !pass {
		t.Errorf("SETTINGS did not match the clone - config file improperly parsed")
	}

}
*/
