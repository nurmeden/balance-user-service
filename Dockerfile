FROM golang:latest as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o app .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/app /app/app

WORKDIR /app

CMD ["./app"]