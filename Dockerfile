FROM golang:1 AS builder
WORKDIR /root/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o dv .

FROM alpine:latest  
RUN apk -U upgrade && \
    apk add ca-certificates 
RUN addgroup -S appgroup && \
    adduser -S appuser -G appgroup
USER appuser
WORKDIR /appuser/
COPY --from=builder /root/dv .
CMD ["/appuser/dv"]
