FROM golang:1.15.6-alpine3.12
RUN go install github.com/tcarvi/go-server@latest

FROM alpine:3.12
COPY --from=0 /go/bin/go-server .
COPY public /public
CMD ["./go-server"]