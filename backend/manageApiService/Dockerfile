FROM golang:1.22-alpine AS builder

WORKDIR /build

ARG proxy=https://proxy.golang.org
ENV proxy=$proxy
RUN echo "proxy: $proxy"

ENV GOPROXY=${proxy}

COPY go.mod .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -ldflags "-X google.golang.org/protobuf/reflect/protoregistry.conflictPolicy=warn" -a -o serve ./cmd/

FROM golang:1.22-alpine AS promtail-builder

WORKDIR /build

ARG proxy=https://proxy.golang.org
ENV proxy=$proxy
RUN echo "proxy: $proxy"

ENV GOPROXY=${proxy}

RUN wget https://github.com/grafana/loki/archive/refs/tags/v3.1.1.zip
RUN unzip v3.1.1.zip

WORKDIR /build/loki-3.1.1

RUN go mod download
RUN CGO_ENABLED=0 go build ./clients/cmd/promtail

FROM alpine:3.16 AS runner

WORKDIR /app

RUN mkdir tmp
RUN mkdir logs
RUN mkdir configs

COPY --from=builder /build/serve /app/serve
COPY ./configs /app/configs
COPY ./entrypoint.sh /app/entrypoint.sh
COPY ./promtail.yaml /app/promtail.yaml

COPY --from=promtail-builder /build/loki-3.1.1/promtail /app/promtail

RUN wget https://dl.min.io/client/mc/release/linux-amd64/mc
RUN chmod -R 777 mc

ENTRYPOINT ["./entrypoint.sh"]
