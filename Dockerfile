FROM golang:1.15.6-alpine3.12
COPY main.go /go/src/github.com/tcarvi/go-server/
COPY handler /go/src/github.com/tcarvi/go-server/handler
COPY go.mod /go/src/github.com/tcarvi/go-server/
RUN go env
RUN go install github.com/tcarvi/go-server@latest

FROM alpine:3.12
COPY --from=0 /go/bin/go-server .
COPY public /public
CMD ["./go-server"]