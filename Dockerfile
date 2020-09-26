FROM golang:1.15-alpine3.12 as builder
WORKDIR /go/api
COPY . .
ENV GO111MODULE=on
COPY go.mod ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o ./

FROM golang:1.15 as dev
WORKDIR /go/api
ENV GO111MODULE=auto
RUN go clean -modcache
COPY . .
RUN go get github.com/pilu/fresh
ENTRYPOINT ["fresh"]
