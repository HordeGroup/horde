FROM golang:1.13.5-alpine3.10 AS build

WORKDIR /go/src/horde
COPY . .
RUN go build -mod vendor -o horde -tags="jsoniter" cmd/app/main.go

FROM alpine:3.10
WORKDIR /app

COPY --from=build /go/src/horde/horde .

CMD ./horde