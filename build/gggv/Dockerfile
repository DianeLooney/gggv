FROM diane/gggv-base:latest

WORKDIR /go/src/app
COPY . .
ENV GO111MODULE=on
RUN go get -d -v ./...
RUN go build cmd/daemon/main.go
