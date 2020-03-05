FROM golang:1.13.8-alpine
WORKDIR /data
COPY go.mod server.go ./
RUN go build server.go
FROM  alpine:3.9.5
WORKDIR /example_write_adapter
COPY --from=0 /data/server /example_write_adapter
ENTRYPOINT ["./server"]



