FROM golang:latest

LABEL maintainer="Santiago Orlando <santiagoelias.orlando@gmail.com"

WORKDIR /filestorage

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

EXPOSE 3003

RUN go build main.go

CMD [ "./main" ]