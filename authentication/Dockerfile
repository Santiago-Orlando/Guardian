FROM golang:1.18.2-alpine3.15

LABEL maintainer="Santiago Orlando <santiagoelias.orlando@gmail.com"

WORKDIR /authentication

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

EXPOSE 3002

RUN go build main.go

CMD [ "./main" ]