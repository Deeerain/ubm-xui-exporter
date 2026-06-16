FROM golang:1.26-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o exporter .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/exporter .

ENV LISTEN_ADDRESS=:2112
ENV METRICS_PATH=/metrics
ENV LOG_LEVEL=info

EXPOSE 2112

CMD [ "./exporter" ]