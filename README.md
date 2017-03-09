# LB-Toggle

- Provides a monitorable toggle switch for load balanced applications, especially those with health and smoke endpoints.

- Enables Operations to easy enable and disable a web application with simple http control endpoints.

- Provides a common format/interface for managing applications which are behind some form of reverse proxy or load balancer that checks at interval for server health status.

## Building

``` bash
git clone https://github.com/vacoj/lb-toggle.git
cd lb-toggle/src/lb-toggle
go build
```

## Configuration file

- To start with a specific configuration file, just execute like this

``` bash
./lb-toggle /path/to/config.json
```

``` json
{
    "targets": {
        "health_endpoint": "http://localhost:5704/fakehealth",  // url to your application's health endpoint
        "smoke_endpoint": "http://localhost:5704/fakesmoke"  // url to your application's smoke endpoint
    },
    "service": {
        "http_port": "5704",  // port to listen on for incoming web requests
        "smoke_interval": 300,  // polling interval for smoke endpoint, in seconds
        "health_interval": 15  // polling interval for health endpoint, in seconds
    }
}
```

## API

### "/status"

- Shows long-form status to the caller

``` json
{
    "State": false,
    "HealthStatus": {
        "OK": true,
        "Last": "2017-03-09T12:35:35.24445478-08:00",
        "Endpoint": "http://localhost:5704/fakehealth",
        "Interval": 15
    },
    "SmokeStatus": {
        "OK": true,
        "Last": "2017-03-09T12:35:35.244431726-08:00",
        "Endpoint": "http://localhost:5704/fakesmoke",
        "Interval": 300
    }
}
```

### "/status/simple"

- Returns 500 error if "State", "HealthStatus.OK", or "SmokeStatus.OK" return false.
- Returns 200 if they all return true.

### "/status/simple2"

- Returns 200 status when "State", "HealthStatus.OK", and "SmokeStatus.OK" are true.
- Returns *NO RESPONSE* error if "State", "HealthStatus.OK", or "SmokeStatus.OK" return false.  This is for F5's health monitor.

### "/toggle/on"

- If "HealthStatus.OK" and "SmokeStatus.OK" are true, sets "State" to true.

### "/toggle/off"

- Sets "State" to false

### "/toggle"

- If "HealthStatus.OK" and "SmokeStatus.OK" are true, sets "State" to true if "!State".
- If "State" is true, then sets "State" to false.

### "/fakesmoke" and "/fakehealth"

- Fake endpoints that return 200, and are the defaults in the config file.  This is just for POC and testing, but if you don't have a smoke endpoint and do have a health endpoint, leave it to what it's already set to in the config file.
