# Builder
FROM golang:1.14.2-alpine3.11

RUN apk update && apk upgrade && \
    apk --update add git make

WORKDIR /app

COPY . .

# Distribution
FROM alpine:latest

RUN apk update && apk upgrade && \
    apk --update --no-cache add tzdata && \
    mkdir /app 

WORKDIR /app 

EXPOSE 8080

COPY --from=builder /app/engine /app

CMD /app/engine