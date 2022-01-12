FROM golang:1.17 as builder

# first (build) stage

WORKDIR /app
COPY main.go .
COPY go.mod .
RUN go mod tidy
RUN go build -mod=mod -o http-server

# final (target) stage

FROM alpine:3.10
COPY --from=builder /app/http-server /
ENTRYPOINT ["/http-server", "--v","2","--logtostderr","true"]
