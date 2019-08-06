FROM diane/gvd-base:latest

WORKDIR /go/src/app
COPY . .
ENV GO111MODULE=on
RUN go get -d -v ./...
RUN go build cmd/main.go
