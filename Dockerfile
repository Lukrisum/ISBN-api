FROM golang:1.18 as builder

RUN go env -w GO111MODULE=auto \
    && go env -w GOPROXY=https://goproxy.cn \
    && go env -w CGO_ENABLED=0

WORKDIR /build

COPY ./go.mod ./go.sum ./

RUN go mod download

COPY ./ ./

RUN go build -ldflags '-extldflags "-static"' -o main

FROM alpine:latest

COPY --from=builder build/main /usr/bin/main

ENTRYPOINT [ "/usr/bin/main" ]