FROM golang:1.17-alpine as builder

RUN apk update && apk add --no-cache git

WORKDIR $GOPATH/src/mypackage/myapp/

COPY . .

RUN go mod download
RUN go mod verify
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -a -installsuffix cgo -o /go/bin/pingpong

FROM scratch

COPY --from=builder /go/bin/pingpong /go/bin/pingpong

ENTRYPOINT ["/go/bin/pingpong"]