## Build
FROM golang:1.19 AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY *.go ./

RUN go build -o goquotemonitor .

## Deploy
# env vars are missing need DATABASE_URL and TOKEN
# run with --env-file ./env to include env vars in runtime

FROM ubuntu:latest
ARG DEBIAN_FRONTEND=noninteractive

RUN apt-get update && apt-get install -y \
  ca-certificates \
  && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY --from=build /app/goquotemonitor .

# TODO: for future web server
EXPOSE 8080

CMD ["./goquotemonitor"]
