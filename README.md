# LB-Health-Toggle

- Provides a monitorable toggle switch for load balanced applications, especially those with health and smoke endpoints.

- Enables Operations to easy enable and disable a web application with simple http control endpoints.

- Provides a common format/interface for managing applications which are behind some form of reverse proxy or load balancer that checks at interval for server health status.

## API

### "/status":  Shows long-form status to the caller - structured as follows:

``` json
{
    "State": false,
    "HealthStatus": {
        "OK": true,
        "Last": "2017-03-08T15:13:46.057289293-08:00"
    },
    "SmokeStatus": {
        "OK": true,
        "Last": "2017-03-08T15:16:16.23510043-08:00"
    }
}
```

### "/status/simple"

- Returns 500 error if "State", "HealthStatus.OK", or "SmokeStatus.OK" return false.
- Returns 200 if they all return true.

### "/toggle/on"

- If "HealthStatus.OK" and "SmokeStatus.OK" are true, sets "State" to true.

### "/toggle/off"

- Sets "State" to false

### "/toggle"

- If "HealthStatus.OK" and "SmokeStatus.OK" are true, sets "State" to true if "!State".
- If "State" is true, then sets "State" to false.


### "/fakesmoke" and "/fakehealth"

- Fake endpoints that return 200, and are the defaults in the config file.  This is just for POC and testing, buf if you don't have a smoke endpoint but do have a health endpoint, leave it to what it's already set to in the config file.
