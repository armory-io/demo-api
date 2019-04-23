FROM golang:1.11-alpine as builder

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOARCH amd64

WORKDIR /go/src/github.com/armory-io/demo-api
COPY . .

RUN go install

FROM alpine:3.8

EXPOSE 3000

WORKDIR /usr/local/bin

RUN apk add -U ca-certificates

COPY --from=builder /go/bin/demo-api .

ENTRYPOINT ["/usr/local/bin/demo-api"]
