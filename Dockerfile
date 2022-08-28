FROM golang:1.19-alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
COPY ./etc /app/etc
RUN go build -ldflags="-s -w" -o /app/5feet11 .


FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Etc/UTC /usr/share/zoneinfo/Etc/UTC
ENV TZ Etc/UTC

WORKDIR /app
COPY --from=builder /app/5feet11 /app/5feet11
COPY --from=builder /app/etc /app/etc

CMD ["./5feet11", "-f", "etc/fivefeeteleven-api.yaml"]
