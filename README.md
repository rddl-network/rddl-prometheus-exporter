# rddl-prometheus-exporter
This service is responsible for providing vital network metrics for prometheus.

## Execution
The service can be executed via the following go command without having it previously built:
```
go run cmd/rddl-prometheus-exporter/*.go
```

## Configuration
The service needs to be configured via the ```./app.toml``` file or environment variables. The defaults are:
```
rpc-host = 'localhost:18884'              // elementsd rpc host
rpc-pass = 'password'                     // elementsd rpc password
rpc-user = 'user'                         // elementsd rpc user
service-bind = 'localhost'
service-port = 8080
service-units = 'foo.service,bar.service' // systemd units
```