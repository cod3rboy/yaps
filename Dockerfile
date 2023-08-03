FROM golang:1.19.1 AS build

## Building Server ##
WORKDIR /src

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o yaps .

FROM debian:stable-20220912-slim AS release

## Running Server ##
WORKDIR /app

COPY --from=build /src/yaps .

RUN apt-get update \
  && apt-get install -y ca-certificates \
  && apt-get clean \
  && chmod +x /app/yaps

EXPOSE 8080

ENTRYPOINT [ "/app/yaps", "-hostName", "0.0.0.0", "-hostPort", $PORT, "-pathPrefix", $PATH_PREFIX ]