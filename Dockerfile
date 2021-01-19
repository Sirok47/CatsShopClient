FROM golang:latest

ENV GO111MODULE=on

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

FROM alpine:latest

COPY --from=0 /app /app

EXPOSE 1323

ENTRYPOINT ["/app/CatsShopClient"]