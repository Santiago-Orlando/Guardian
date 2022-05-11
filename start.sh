#!/bin/bash


host_IP=$(hostname -I | awk '{split($0,a," "); print a[1]}')

sed -i '23 i ENV MONGO_URI="mongodb://'"$host_IP"':3005/authentication"' authentication/Dockerfile 

sed -i '13 i ENV AUTHENTICATION_URL=''"http://'"$host_IP"'"' errorLogger/Dockerfile
sed -i '14 i ENV FILESTORAGE_URL=''"http://'"$host_IP"'"' errorLogger/Dockerfile
sed -i '18 i ENV POSTGRESQL_URI=''"postgresql://@'"$host_IP"':3007"' errorLogger/Dockerfile

sed -i '20 i ENV POSTGRESQL_URI=''"postgresql://@'"$host_IP"':3006"' fileStorage/Dockerfile

sed -i '17 i ENV AUTHENTICATION_URL=''"http://'"$host_IP"'"' proxy/Dockerfile
sed -i '18 i ENV PORT_AUTHENTICATION_SERVICE="3002"' proxy/Dockerfile
sed -i '19 i ENV FILESTORAGE_URL=''"http://'"$host_IP"'"' proxy/Dockerfile
sed -i '20 i ENV PORT_FILESTORAGE_SERVICE="3003"' proxy/Dockerfile

docker pull mongo
docker pull postgres

docker create --name mongo_db -p 3005:27017 mongo

docker build -t error_logger ./errorLogger/
docker build -t error_logger_db ./errorLogger/database/
docker build -t authentication ./authentication/ 
docker build -t file_storage ./fileStorage/
docker build -t file_storage_db ./fileStorage/api/database/
docker build -t proxy ./proxy/

docker create --name error_logger -p 3001:3001 error_logger
docker create --name error_logger_db -p 3007:3007 error_logger_db
docker create --name authentication -p 3002:3002 authentication
docker create --name file_storage -p 3003:3003 file_storage
docker create --name file_storage_db -p 3006:3006 file_storage_db
docker create --name proxy -p 3004:3004 proxy

docker start mongo_db 
docker start error_logger_db
docker start file_storage_db

docker start error_logger 
docker start authentication 
docker start file_storage 
docker start proxy 

echo ""
echo "Press any key to stop the services!"
read end

docker kill $(docker ps -q)

echo "Bye!"
exit 0
