FROM golang:1.13-alpine3.11 AS build
RUN apk --no-cache add gcc g++ make ca-certificates
WORKDIR /go/src/github.com/sudatra/go-microservice
COPY go.mod g.sum ./
COPY vendor vendor
COPY account account
COPY catalog catalog
COPY order order

RUN GO111MODULE=on go build -mod vendor -o /go/bin/app ./account/cmd/order

FROM alpine:3.11
WORKDIR /usr/bin
COPY --from=build /go/bin .
EXPOSE 8080
CMD ["app"]