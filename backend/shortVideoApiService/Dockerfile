FROM golang:1.22-alpine AS builder

WORKDIR /build

ARG proxy=https://proxy.golang.org
ENV proxy=$proxy
RUN echo "proxy: $proxy"

ENV GOPROXY ${proxy}

COPY go.mod .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -ldflags "-X google.golang.org/protobuf/reflect/protoregistry.conflictPolicy=warn" -a -o serve ./cmd/

FROM alpine:3.16 AS runner

WORKDIR /app

RUN mkdir tmp
RUN mkdir logs
RUN mkdir configs

COPY --from=builder /build/serve /app/serve
COPY ./configs /app/configs

ENTRYPOINT ["/app/serve"]