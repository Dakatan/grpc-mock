FROM golang:1.15 AS builder
WORKDIR /build
COPY go.sum go.sum
COPY go.mod go.mod
RUN go mod download
ARG build_target
COPY . .
RUN make $build_target

FROM alpine:3.12
RUN apk add --no-cache tzdata
ENV TZ=Asia/Tokyo
COPY --from=builder /build/app /usr/local/bin/app
EXPOSE 50051
ENTRYPOINT ["/usr/local/bin/app"]
