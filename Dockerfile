FROM golang:alpine3.18 AS build
WORKDIR /build
COPY . .
ENV CGO_ENABLED=1
ENV GOOS=linux
RUN apk add --no-cache \
    gcc \
    musl-dev

RUN go mod download
RUN go build -o ./bin/main ./cmd/main.go

FROM alpine:3 AS final
WORKDIR /app
RUN apk --update upgrade && \
    apk add sqlite && \
    rm -rf /var/cache/apk/*
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
COPY --from=build /build/bin/main ./main
RUN chmod +x ./main
CMD ["./main"]