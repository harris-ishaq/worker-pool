## Builder
FROM golang:1.17-alpine3.15 as builder

LABEL name="batch-migcontract"
LABEL version="1.0.0"

RUN apk update && apk add --no-cache git
RUN apk --no-cache add tzdata
RUN apk --no-cache add ca-certificates

RUN mkdir -p /home/go/app

WORKDIR /home/go/app

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o housekeep

## Distribution
FROM scratch

COPY --from=builder /home/go/app/housekeep /home/go/app/server
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY ./credential.json /home/go/app/credential.json
ENV TZ=Asia/Jakarta

ENTRYPOINT  ["/home/go/app/server"]