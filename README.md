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

``` javascript
{
    "targets": [
        {
            "endpoint": "http://localhost:5704/fakehealth",  // url to your application's health endpoint
            "polling_interval": 15,  // polling interval for target endpoint, in seconds
            "expected_status_code": 200,  // HTTP status code to look for.  If this isn't returned when the check happens, we mark OK as false.
        },
        {
            "name": "FakeSmoke"  // Arbitrary - use for your own reasons, or leave it blank.
            "endpoint": "http://localhost:5704/fakesmoke",  // url to your application's health endpoint
            "polling_interval": 300,  // polling interval for target endpoint, in seconds
            "expected_status_code": 200,  // HTTP status code to look for.  If this isn't returned when the check happens, we mark OK as false.
            "expected_response_string": ""  // response is parsed for this string.  If expected_response_string is blank, check is ignored.  If found, OK is true
            "unexpected_response_string": ""  // response is parsed for this string.  If unexpected_response_string is blank, check is ignored.  If found, OK is false  (an example would be searching repsonse text for {"thisthing": false}, and if found, causes OK to be set to false)
        }
    ],
    "service": {
        "http_port": "5704",  // *string not int; port to listen on for incoming web requests
    }
}
```

## API

### "/status"

- Shows long-form status to the caller

``` javascript
{
    "State": false,  // If both Health and Smoke are OK, this can be enabled.  If either of the mentioned are *NOT OK*, will be false.  If both endpoints are OK, this can be toggled true and false.
    "Targets": [{
        "name": "FakeHealth"
        "OK": true,  // true if OK, false if BAD
        "Last": "2017-03-09T12:35:35.24445478-08:00",  // last time the status of the health endpoint was OK
        "Endpoint": "http://localhost:5704/fakehealth",  // URL to check for status
        "Interval": 15  // interval in seconds between polls
    },
    {
        "Name": "FakeSmoke"  // Arbitrary - use for your own reasons, or leave it blank.
        "OK": true,  // true if OK, false if BAD
        "Last": "2017-03-09T12:35:35.24445478-08:00",  // last time the status of the health endpoint was OK
        "Endpoint": "http://pirri.vacovsky.us/login",  // URL to check for status
        "Interval": 300  // interval in seconds between polls
    }],
    "Version": 0.2.0  // version of the health check app
}
```

### "/status/simple"

- Returns 503 (Service Unavailable) error if "State", "HealthStatus.OK", or "SmokeStatus.OK" return false.
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
