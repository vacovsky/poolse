package main

//Settings contains the config.json information for configuring the listening port, monitored application details, etc
type Settings struct {
	Target struct {
		HealthEndpoint string // //"health_endpoint": "http://localhost/health",
	}
	Service struct {
		HTTPPort int //5704,
	}
}

func (s *Settings) parseSettingsFile() {

}
