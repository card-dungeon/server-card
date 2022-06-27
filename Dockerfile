FROM golang:1.18.3-alpine AS builder
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

# linux
RUN apk add --no-cache upx 
# OR
# RUN apt update && apt install -y xz-utils && rm -rf /var/lib/apt/lists/*
# ADD https://github.com/upx/upx/releases/download/v3.95/upx-3.95-amd64_linux.tar.xz /usr/local
# RUN xz -d -c /usr/local/upx-3.95-amd64_linux.tar.xz | tar -xOf - upx-3.95-amd64_linux/upx > /bin/upx && chmod a+x /bin/upx

WORKDIR /build

COPY go.mod .
COPY go.sum .

RUN go mod tidy
COPY . .
RUN go build -a -ldflags '-s -w' -o main main.go
RUN upx --lzma --best main

FROM scratch
COPY --from=builder /build .
EXPOSE 10503
CMD ["/main"]