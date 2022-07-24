FROM golang:1.18.1-alpine as build
WORKDIR /app
# Required for VCS stamping.
RUN apk add --no-cache git
COPY . .
RUN go build .

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/api-gateway .
ENTRYPOINT [ "./api-gateway" ]
