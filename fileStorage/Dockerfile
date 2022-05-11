FROM golang:latest

LABEL maintainer="Santiago Orlando <santiagoelias.orlando@gmail.com"

WORKDIR /filestorage

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .


ENV PORT_FILESTORAGE_SERVICE="3003"

ENV JWTSecret="secret"

ENV POSTGRESQL_URI="postgresql://@192.168.1.134:3006"
ENV POSTGRESQL_URI="postgresql://@192.168.1.134:3006"



RUN go build main.go

CMD [ "./main" ]