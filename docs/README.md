# API Gateway

The api-gateway is a component of the Noted backend which is responsible for proxying HTTP requests to the Noted microservice backend. It receives requests in `application/json` format, forwards them to the corresponding gRPC service and returns the response. 

## Configuration

| Env Name                      | Flag Name                 | Default         | Description                               |
|-------------------------------|---------------------------|-----------------|-------------------------------------------|
| `APIGW_PORT`                  | `--port`                  | `3000`          | The port the application shall listen on. |
| `APIGW_ENV`                   | `--env`                   | `production`    | Either `production` or `development`.     |
| `APIGW_ACCOUNTS_SERVICE_ADDR` | `--accounts-service-addr` | `accounts:3000` | The gRPC address of the accounts service. |
