FROM golang:latest

LABEL maintainer="Santiago Orlando <santiagoelias.orlando@gmail.com"

WORKDIR /authentication

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .


ENV PORT_AUTHENTICATION_SERVICE="3002"

ENV GMAIL=""
ENV GMAIL_PW=""

ENV JWTSecret="secret"

ENV MONGO_URI="mongodb://192.168.1.134:3005/authentication"
ENV MONGO_URI="mongodb://192.168.1.134:3005/authentication"

RUN go build main.go

CMD [ "./main" ]