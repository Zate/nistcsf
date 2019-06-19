FROM golang:latest as builder
WORKDIR /go/src/github.com/zate/simplenist/

COPY main.go .
COPY favicon.ico .
COPY common.css .
COPY public public
COPY static static
RUN dep init && dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o simplenist .

FROM scratch
WORKDIR /
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /go/src/github.com/zate/simplenist/simplenist .
COPY --from=builder /go/src/github.com/zate/simplenist/common.css .
COPY --from=builder /go/src/github.com/zate/simplenist/favicon.ico .
COPY --from=builder /go/src/github.com/zate/simplenist/public /public
COPY --from=builder /go/src/github.com/zate/simplenist/static /static
EXPOSE 2086
CMD ["/simplenist"]