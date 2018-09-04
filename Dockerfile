FROM golang:latest
WORKDIR /go/src/github.com/dblooman/dblooman-go-app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build .

FROM alpine:latest
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

COPY --from=0 /go/src/github.com/dblooman/dblooman-go-app/dblooman-go-app .

EXPOSE 8080

CMD ["./dblooman-go-app"]