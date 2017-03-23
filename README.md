# Go-Healthcheck

- Provides a monitorable application healthcheck & status endpoint, as well as a "toggle switch" for soft upping and downing of monitored applications.

- Enables automation and operations personnel to easily enable and disable an upstream node/application with simple http control endpoints.

- JSON formatted status results allow for easy status monitoring of target applications from external monitoring tools.  Can be on a different system than the monitored application.

- Provides a common format/interface for managing applications which are behind some form of reverse proxy or load balancer that checks at interval for server health status.

## Building

``` bash
go get -u github.com/davecgh/go-spew/spew
git clone https://github.com/vacoj/go-healthcheck.git
cd go-healthcheck/src/go-healthcheck
go build
```

## Configuration file

- To start with a specific configuration file, just execute like this

``` bash
./go-healthcheck /path/to/config.json
```

``` javascript
{
    "state": {
        "startup_state": true, // if true, when server starts, if first check passes server state is marked OK
        "administrative_state": "AdminOff",  // If persistent state isn't on, this is the default statup state for the STATUS.  If will only be OK if all Targets are also OK on first check
        "persist_state": true  // indicates whether or not STATUS.State.AdministrativeState should be sticky between settings/application restarts and reloads.
    },
    "targets": [
        {
            "endpoint": "http://localhost:5704/fakehealth",  // url to your application's health endpoint
            "polling_interval": 15,  // polling interval for target endpoint, in seconds
            "expected_status_code": 200,  // *required* HTTP status code to look for.  If this isn't returned when the check happens, we mark OK as false.
            "up_count_threshold": 10, // this many failed checks will mark target as online
            "down_count_threshold": 10 // this many failed checks will mark target as offline
        },
        {
            "name": "Expected Example",
            "endpoint": "http://localhost:5704/fakeexpected",  // url to your application's health endpoint
            "polling_interval": 20,  // polling interval for target endpoint, in seconds
            "expected_status_code": 200,  // HTTP status code to look for.  If this isn't returned when the check happens, we mark OK as false.
            "expected_response_strings": ["{\"is_working\": true}"],
            "up_count_threshold": 10, // this many failed checks will mark target as online
            "down_count_threshold": 10 // this many failed checks will mark target as offline
        },
        {
            "name": "Unexpected Example",
            "endpoint": "http://localhost:5704/fakeexpected",  // url to your application's health endpoint
            "polling_interval": 10,  // polling interval for target endpoint, in seconds
            "expected_status_code": 200,  // HTTP status code to look for.  If this isn't returned when the check happens, we mark OK as false.
            "unexpected_response_strings": ["{\"is_working\": false}"], // response is parsed for this string.  If unexpected_response_string is blank, check is ignored.  If found, OK is false  (an example would be searching repsonse text for {"thisthing": false}, and if found, causes OK to be set to false)
            "up_count_threshold": 10, // this many failed checks will mark target as online
            "down_count_threshold": 10 // this many failed checks will mark target as offline
        },
        {
            "name": "Fake Smoke"  // Arbitrary - use for your own reasons, or leave it blank.
            "endpoint": "http://localhost:5704/fakesmoke",  // url to your application's health endpoint
            "polling_interval": 300,  // polling interval for target endpoint, in seconds
            "expected_status_code": 200,  // *required* HTTP status code to look for.  If this isn't returned when the check happens, we mark OK as false.
        }
    ],
    "service": {
        "http_port": "5704",  // *string not int; port to listen on for incoming web requests
        "debug": false, // displays certain pieces of data in console if true
        "show_http_log": true // shows log in console of calls being made if true
    }
}
```

## API

### Help Endpoints

#### "/"

- Displays README.md to caller

### Status Endpoints

#### "/status"

- Shows long-form status to the caller
- All requests to endpoints stemming from "/status" will accept a query string param of ```id``` (example: ```/status/simple?id=1```) to get just the target object at that index; id should correspond to index of target in config file.

``` javascript
{
    "State": {
        "ok": true,
        "startup_state": true,
        "persist_state": true,
        "administrative_state": ""
    },
    "Targets": [
        {
            "id": 0,
            "name": "",
            "endpoint": "http://localhost",
            "polling_interval": 5,
            "expected_status_code": 200,
            "expected_response_strings": null,
            "unexpected_response_strings": null,
            "last_ok": "2017-03-23T08:30:51.1019184-07:00",
            "last_checked": "2017-03-23T08:30:51.1019184-07:00",
            "ok": true,
            "up_count": 6987,
            "up_count_threshold": 10,
            "down_count": 0,
            "down_count_threshold": 1
        }
    ],
    "Version": "0.3.10"
}
```

#### "/status/simple"

- Returns 200 if they all return true.
- Returns *NO RESPONSE* error if any of the Targets (targets[i].ok) are false, or if State is false.  This is for F5's health monitor.

#### "/status/simple2"

- Returns 200 status when "State", and all of the Targets (targets[i].ok) are true.
- Returns *NO RESPONSE* error if any of the Targets (targets[i].ok) are false, or if State is false.  This is for F5's health monitor.

### Toggle Endpoints

#### "/toggle"

- If "HealthStatus.OK" and "SmokeStatus.OK" are true, sets "State" to true if "!State".
- If "State" is true, then sets "State" to false.

#### "/toggle/on"

- If "HealthStatus.OK" and "SmokeStatus.OK" are true, sets "State" to true.

#### "/toggle/off"

- Sets "State" to false

#### "/toggle/adminoff"

- Prevents happy status no matter what else is "OK"

#### "/toggle/adminon"

- Prevents unsuccessful response no matter what else is not "OK"

#### "/toggle/adminreset"

- Clears state.dat and resets AdministrativeState to empty

### Settings Endpoints

#### "/settings"

- Displays currently loaded Settings struct, as populated from the config file.

#### "/settings/reload"

- Reloads configuration file into memory and restarts all target monitors contained within.  A configuration reload generally takes as much time as the longest polling interval present in the targets list, plus 5 seconds.

### Testing Endpoints

#### "/fakesmoke", "fakeexpected", and "/fakehealth"

- Fake endpoints that return 200, maybe some context, and are the defaults in the config file.  This is just for POC and testing, but if you don't have a smoke endpoint and do have a health endpoint, leave it to what it's already set to in the config file, or delete the targets from the config.
