FROM golang:1.17 as builder

# first (build) stage

ENV GO111MODULE=on
WORKDIR /app
COPY main.go .
COPY pkg .
COPY go.mod .
RUN go mod download
RUN go build -mod=mod -o http-server

# final (target) stage

FROM scratch
COPY --from=builder /app/http-server /
ENTRYPOINT ["/http-server", "--v","2","--logtostderr","true"]
