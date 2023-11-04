FROM golang:1.21-alpine
COPY main.go /go/src/github.com/tcarvi/go-server/
COPY handler /go/src/github.com/tcarvi/go-server/handler
COPY go.mod /go/src/github.com/tcarvi/go-server/
RUN go install github.com/tcarvi/go-server@latest

FROM alpine:latest
COPY --from=0 /go/bin/go-server .
COPY public /public
CMD ["./go-server"]