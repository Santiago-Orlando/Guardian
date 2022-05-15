FROM node:latest

LABEL maintainer="Santiago Orlando <santiagoelias.orlando@gmail.com"

WORKDIR /errorLogger

COPY . .

RUN npm i

EXPOSE 3001

CMD [ "node", "server.js" ]