#!/bin/bash

chmod 777 config.sh
source config.sh


cd authentication
go build main.go

cd ../fileStorage
go build main.go

cd ../proxy
go build main.go

cd ..


createdb "Guardian_Files" > /dev/null 2>&1

mkdir fileStorage/uploads > /dev/null 2>&1


function Error_Logger() {

    npm i --prefix errorLogger/
    npm start --prefix errorLogger/ 

}

function Load_Services(){

    ./authentication/main &
    ./fileStorage/main &
    ./proxy/main &
    Error_Logger &

}

function Services_Killer() {

    lsof -i tcp":$PORT_ERROR_SERVICE" | awk "/${PORT_AUTHENTICATION_SERVICE:1}"'/{print $2}' | xargs kill > /dev/null 2>&1
    lsof -i tcp":$PORT_AUTHENTICATION_SERVICE" | awk "/${PORT_AUTHENTICATION_SERVICE:1}"'/{print $2}' | xargs kill > /dev/null 2>&1
    lsof -i tcp":$PORT_FILESTORAGE_SERVICE" | awk "/${PORT_FILESTORAGE_SERVICE:1}"'/{print $2}' | xargs kill > /dev/null 2>&1
    lsof -i tcp":$PORT_PROXY_SERVICE" | awk "/${PORT_PROXY_SERVICE:1}"'/{print $2}' | xargs kill > /dev/null 2>&1

} 


createdb "Guardian_Errors" > /dev/null 2>&1


psql $USER -h $POSTGRESQL_URL -d "Guardian_Errors" -f errorLogger/database/TABLES.SQL > /dev/null 2>&1 &
psql $USER -h $POSTGRESQL_URL -d "Guardian_Files" -f fileStorage/api/database/FILES.SQL > /dev/null 2>&1 &


Load_Services > /dev/null 2>&1 &
open http://localhost:${PORT_PROXY_SERVICE}/docs > /dev/null 2>&1 &


echo "If you want to stop the services press any letter."

read end

Services_Killer > /dev/null 2>&1

echo "Bye!"
exit 0