#
# STAGE 1: prepare
#
FROM golang:1.18-alpine as prepare

WORKDIR /source

COPY vendor .

#
# STAGE 2: build
#
FROM prepare AS build

COPY . .

RUN CGO_ENABLED=0 go build -o bin/kubectldb cmd/app/main.go

#
# STAGE 3: run
#
FROM scratch as run
USER 10001:10001

COPY --from=build /source/bin/kubectldb /kubectldb

EXPOSE 8080

ENTRYPOINT ["/kubectldb"]