#
# 1. Build Container
#
FROM golang:1.20-alpine AS build

ENV GO111MODULE=on \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

RUN mkdir -p /app

COPY go.sum go.mod /app/

RUN go mod download

COPY . /app

RUN go build -o ./dist/private_banking /app/init/main.go

#
# 2. Runtime Container
#
FROM alpine:3.17 as production

LABEL maintainer="Mahdi Imani <imani.mahdi@gmail.com>"

ENV TZ=Asia/Shanghai

RUN apk update && \
    apk add --update --no-cache wget tzdata && \
    cp --remove-destination /usr/share/zoneinfo/${TZ} /etc/localtime && \
    echo "${TZ}" > /etc/timezone

WORKDIR /app

COPY --from=build /app/dist /app/
RUN mkdir /app/logs

EXPOSE 3000

CMD ["./private_banking", "d"]
