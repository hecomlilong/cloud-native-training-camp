FROM golang:1.17 as builder

# first (build) stage

ENV GO111MODULE=on
WORKDIR /app
COPY main.go .
COPY pkg .
COPY go.mod .
RUN go mod tidy
RUN go build -o http-server

# final (target) stage

FROM alpine
COPY --from=builder /app/http-server /
RUN apk add libc6-compat
ENTRYPOINT ["/http-server", "--v","2","--logtostderr","true"]
