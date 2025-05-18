FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o cinehive-be main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/cinehive-be .
EXPOSE 8000
CMD ["./cinehive-be"]