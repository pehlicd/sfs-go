# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY files/ ./files/
COPY go.mod ./
COPY *.go ./

RUN go build -o /sfs-go

EXPOSE 4000

CMD [ "/sfs-go" ]