#!/bin/bash


export PORT_ERROR_SERVICE="3001"
export PORT_AUTHENTICATION_SERVICE=":3002"
export PORT_FILESTORAGE_SERVICE=":3003"

export MONGO_URI="mongodb://127.0.0.1:27017/test"
export POSTGRESQL_URI="postgresql://@127.0.0.1:5432/"
export POSTGRESQL_URL="127.0.0.1"

export JWTSecret="SuperDifficultPassword"

GMAIL=""
GMAIL_PW=""
ADMIN_EMAIL=""


cd authentication
go build main.go

cd ../fileStorage
go build main.go

cd ..


createdb "Guardian_Files" > /dev/null 2>&1

mkdir fileStorage/uploads > /dev/null 2>&1


function Error_Logger() {

    npm i --prefix errorLogger/
    npm start --prefix errorLogger/ 

}

function Base_Services(){

    psql $USER -h $POSTGRESQL_URL -d "Guardian_Files" -f fileStorage/api/database/FILES.SQL > /dev/null 2>&1 &

    ./authentication/main &
    ./fileStorage/main &

    open http://localhost:${PORT_AUTHENTICATION_SERVICE:1}/docs &
    open http://localhost:${PORT_FILESTORAGE_SERVICE:1}/docs &

}

function Services_Killer() {

    lsof -i tcp:$PORT_ERROR_SERVICE | awk "/${PORT_AUTHENTICATION_SERVICE:1}"'/{print $2}' | xargs kill > /dev/null 2>&1
    lsof -i tcp$PORT_AUTHENTICATION_SERVICE | awk "/${PORT_AUTHENTICATION_SERVICE:1}"'/{print $2}' | xargs kill > /dev/null 2>&1
    lsof -i tcp$PORT_FILESTORAGE_SERVICE | awk "/${PORT_FILESTORAGE_SERVICE:1}"'/{print $2}' | xargs kill > /dev/null 2>&1

} 


if [ "$1" == "--full" ]; then
    read -p "Please insert the error service port: " PORT_ERROR_SERVICE
    read -p "Please insert the authentication service port: " PORT_AUTHENTICATION_SERVICE
    read -p "Please insert the file storage service port: " PORT_FILESTORAGE_SERVICE

    read -p "Please insert the mongo URI: " MONGO_URI
    read -p "Please insert the postgreSQL URI: " POSTGRESQL_URI
    read -p "Please insert the postgreSQL URL: " POSTGRESQL_URL
    read -p "Please insert the postgreSQL password: " POSTGRESQL_PASSWORD

    read -p "Please insert the secret to generate the tokens: " JWTSecret

    read -p "Please insert the gmail of the app: " GMAIL
    read -p "Please insert the gmail aplication password: " GMAIL_PW
    read -p "Please insert the admin email: " ADMIN_EMAIL

    PORT_AUTHENTICATION_SERVICE=":$PORT_AUTHENTICATION_SERVICE" 
    PORT_FILESTORAGE_SERVICE=":$PORT_FILESTORAGE_SERVICE"

    export PORT_ERROR_SERVICE
    export PORT_AUTHENTICATION_SERVICE
    export PORT_FILESTORAGE_SERVICE

    export MONGO_URI
    export POSTGRESQL_URI
    export POSTGRESQL_PASSWORD

    export JWTSecret

    export GMAIL
    export GMAIL_PW
    export ADMIN_EMAIL

else

    Base_Services > /dev/null 2>&1 & 

    echo "If you want to stop the services press any letter."

    read end

    Services_Killer > /dev/null 2>&1 &
    
    echo "Bye!"
    exit 0

fi

createdb "Guardian_Errors" > /dev/null 2>&1


psql $USER -h $POSTGRESQL_URL -d "Guardian_Errors" -f errorLogger/database/TABLES.SQL > /dev/null 2>&1 &


Error_Logger > /dev/null 2>&1 &
Base_Services > /dev/null 2>&1 &

echo "If you want to stop the services press any letter."

read end

Services_Killer > /dev/null 2>&1

echo "Bye!"
exit 0