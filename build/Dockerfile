FROM golang:1.12.0-alpine3.9 as build

ENV CGO_ENABLED="0" GOOS="linux" GOARCH="amd64"

WORKDIR /build

ADD .. /build

RUN go build -o ports .

FROM scratch
USER 10001:10001

COPY --from=build /ports /ports

ENTRYPOINT ["/ports"]