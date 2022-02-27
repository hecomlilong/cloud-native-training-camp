FROM golang:1.17 as builder

# first (build) stage

ENV GO111MODULE=on
ENV GOMOD=/app/go.mod
WORKDIR /app
COPY mod2 .
COPY pkg .
COPY go.mod .
RUN go mod tidy
RUN go build -o http-server mod2/main.go

# final (target) stage

FROM alpine
COPY --from=builder /app/http-server /
RUN apk add libc6-compat
ENTRYPOINT ["/http-server", "--v","2","--logtostderr","true"]
