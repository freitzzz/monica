FROM golang:1.19.3-alpine AS builder
WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY cmd/ ./cmd/
COPY internal/ ./internal/

RUN GOOS=linux CGO_ENABLED=0 go build -o monica ./cmd/monica/monica.go

FROM scratch
WORKDIR /app

COPY --from=builder /app /app/
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

CMD [ "/app/monica" ]