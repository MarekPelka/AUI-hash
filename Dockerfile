FROM golang:1.9.1 as builder
WORKDIR /go/src/AUI-hash/
COPY hasher.go .
COPY hasher_test.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o auiHash .

FROM scratch
WORKDIR /root/
COPY --from=builder /go/src/AUI-hash/auiHash .
CMD ["./auiHash"]
