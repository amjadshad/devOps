FROM golang:1.19.3-alpine3.16 as builder

WORKDIR /go/src/app/
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 go install -installsuffix "static" ./...

FROM alpine:3.16.2 as runtime

RUN set -x \
  && apk add --update --no-cache ca-certificates tzdata \
  && echo 'Etc/UTC' > /etc/timezone \
  && update-ca-certificates

ENV TZ=/etc/localtime   \
    LANG=en_US.utf8     \
    LC_ALL=en_US.UTF-8

COPY --from=builder /go/bin/app /

RUN chmod +x /app && adduser -S appuser -u 1000 -G root
USER 1000

ENTRYPOINT ["/app"]
