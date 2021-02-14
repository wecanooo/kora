FROM golang:1.14-buster as builder

WORKDIR /tmp/tiny-golang-image
COPY . .

RUN go mod tidy \
    && go get -u -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-s -w' -o main

FROM scratch
COPY --from=builder /tmp/tiny-golang-image /
CMD ["/main", "server"]