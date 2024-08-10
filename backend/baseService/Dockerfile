FROM golang:1.22-alpine3.16 AS builder

ARG proxy=https://proxy.golang.org
ENV proxy=$proxy
RUN echo "proxy: $proxy"

ENV GOPROXY ${proxy}

COPY go.mod .
RUN go mod tidy

COPY . .
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -o serve ./cmd/

FROM alpine:3.16 AS runner

WORKDIR /app

RUN mkdir tmp
RUN mkdir logs

COPY --from=builder /build/serve /app/serve
COPY --from=builder /build/configs /app/configs

ENTRYPOINT ["/app/serve"]