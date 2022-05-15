FROM golang:1.18.2-alpine3.15

LABEL maintainer="Santiago Orlando <santiagoelias.orlando@gmail.com"

WORKDIR /proxy

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

EXPOSE 3004

RUN go build main.go

CMD [ "./main" ]