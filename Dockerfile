# Builder
FROM golang:1.15.1-alpine AS builder

RUN apk update && apk add --no-cache git make

WORKDIR /app
COPY . .
RUN make

# Distribution
FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /app/bin/main .

# EXPOSE 8088
CMD ["./main"]  