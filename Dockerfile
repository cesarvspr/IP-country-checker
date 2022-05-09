# syntax=docker/dockerfile:1

FROM golang:1.18-alpine
RUN apk update && apk add git
# Set Environment Variables
ENV HOME /app
ENV CGO_ENABLED 0
ENV GOOS linux
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .

COPY . .

EXPOSE 8013


RUN go build -o /main



CMD [ "/main" ]