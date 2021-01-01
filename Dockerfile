FROM golang:1.15.6-alpine3.12
COPY main.go /go/src/github.com/eduardoleal1981/go-server/
COPY handler /go/src/github.com/eduardoleal1981/go-server/handler
COPY go.mod /go/src/github.com/eduardoleal1981/go-server/
RUN go install github.com/eduardoleal1981/go-server

FROM alpine:3.12
COPY --from=0 /go/bin/go-server .
COPY public /public
EXPOSE 8080
CMD ["./go-server"]