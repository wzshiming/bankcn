FROM golang:alpine AS builder
WORKDIR /go/src/github.com/wzshiming/bankcn/
COPY . .
RUN go install github.com/wzshiming/bankcn/cmd/...


FROM wzshiming/upx AS upx
COPY --from=builder /go/bin/ /go/bin/
RUN upx /go/bin/*

FROM alpine
RUN apk add -U --no-cache ca-certificates openssl tzdata
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY --from=upx /go/bin/ /usr/local/bin/
CMD bankcn
