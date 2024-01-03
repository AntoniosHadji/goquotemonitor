## Build
FROM golang:1.21 AS build

WORKDIR /app

COPY .  ./
RUN go mod download

RUN go build -o goquotemonitor .

## Deploy
# env vars are not included. need DATABASE_URL and TOKEN
# run with --env-file ./env to include env vars in runtime

FROM ubuntu:latest
ARG DEBIAN_FRONTEND=noninteractive

RUN apt-get update && apt-get install -y \
  ca-certificates \
  && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY --from=build /app/goquotemonitor .

# web server UI
EXPOSE 8080

CMD ["./goquotemonitor"]
