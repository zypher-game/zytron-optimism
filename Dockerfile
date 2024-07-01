FROM golang:1.22-bookworm AS builder

WORKDIR /build
ADD . .

RUN make op-node op-batcher op-proposer

FROM debian:bookworm
RUN apt-get update -y
RUN apt-get install -y curl
RUN apt-get install -y ca-certificates

WORKDIR /app
COPY --from=builder /build/op-node/bin/op-node .
COPY --from=builder /build/op-batcher/bin/op-batcher .
COPY --from=builder /build/op-proposer/bin/op-proposer .
