FROM golang:latest AS builder

WORKDIR /app

COPY go.* .
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o processor

FROM alpine:3.12

RUN apk --no-cache add tzdata ca-certificates libc6-compat ffmpeg
COPY --from=builder /app/processor /app/processor

WORKDIR /app
COPY .env .

CMD /app/processor
