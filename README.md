# API Gateway

Service responsible for proxying HTTP requests to the Noted gRPC microservice backend.

## Build

To run the service you only need to have [golang](https://go.dev) and [docker](https://docs.docker.com/get-docker/) installed.

Upon cloning the repository run:
```
make init
```

To generate the protobuf server and stubs:
```
make codegen
```

You can then build the project using the go toolchain.

## Configuration

| Env Name                            | Flag Name                 | Default         | Description                               |
|-------------------------------------|---------------------------|-----------------|-------------------------------------------|
| `API_GATEWAY_PORT`                  | `--port`                  | `3000`          | The port the application shall listen on. |
| `API_GATEWAY_ENV`                   | `--env`                   | `production`    | Either `production` or `development`.     |
| `API_GATEWAY_ACCOUNTS_SERVICE_ADDR` | `--accounts-service-addr` | `accounts:3000` | The gRPC address of the accounts service. |
