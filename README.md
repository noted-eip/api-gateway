# API Gateway

Service responsible for proxying HTTP requests to the Noted gRPC microservice backend.

## Build

To run the service you only need to have [golang](https://go.dev) and [docker](https://docs.docker.com/get-docker/) installed.

Upon cloning the repository run:

```
make update-submodules
```

You can then build the project using the go toolchain.

## Configuration

| Env Name                                   | Flag Name                        | Default                | Description                                      |
|--------------------------------------------|----------------------------------|------------------------|--------------------------------------------------|
| `API_GATEWAY_PORT`                         | `--port`                         | `3000`                 | The port the application shall listen on.        |
| `API_GATEWAY_ENV`                          | `--env`                          | `production`           | Either `production` or `development`.            |
| `API_GATEWAY_ACCOUNTS_SERVICE_ADDR`        | `--accounts-service-addr`        | `accounts:3000`        | The address of the gRPC accounts service.        |
| `API_GATEWAY_NOTES_SERVICE_ADDR` | `--notes-service-addr` | 
`notes:3000` | The address of the gRPC notes service. |
| `API_GATEWAY_RECOMMENDATIONS_SERVICE_ADDR` | `--recommendations-service-addr` | `recommendations:3000` | The address of the gRPC recommendations service. |
