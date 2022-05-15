#!/bin/bash


docker-compose build 

docker-compose up -d

echo ""
echo "Press any key to stop the services!"
read end

docker-compose down

echo "Bye!"
exit 0
