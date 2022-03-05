FROM golang:1.16-alpine as builder

RUN apk update && apk add --no-cache git

WORKDIR $GOPATH/src/mypackage/myapp/

COPY . .

RUN go mod download
RUN go mod verify
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -a -installsuffix cgo -o /go/bin/log

FROM scratch

COPY --from=builder /go/bin/log /go/bin/log

ENTRYPOINT ["go/bin/log"]


