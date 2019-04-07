FROM golang:1.11.5 as builder

ENV GO111MODULE on

WORKDIR /go/src/github.com/kalmeshbhavi/go-examples/goroutines_03

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/kalmeshbhavi/go-examples/goroutines_03/helloworld .

CMD ["./helloworld"]
