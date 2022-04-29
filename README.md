# Guardian

## Proyect description

"Guardian" is a file-saving application and is made up of 3 microservices described below.

- Authentication registers, logs and allows users to change their passwords. It validates all the data and saves them in a __MongoDB__ database, at login it sets a cookie with a __JWT__ token.

- Files Storage Receives, compresses, hashes, and stores files. It stores in a __PostgreSQL__ database the hash of the file (among other data) not allowing the same file to be uploaded twice, even if it belongs to two different users, both can access the file but it is only stored once.

- Error Logger" Receives and stores errors, separated by different failure points, such as "authentication" or "system". This service pings the other two every minute and if any of them does not respond, an email is sent to the administrator. Made it in __JavaScript__

The Authentication and File Storage services are documented using __swagger__.

## Start the Proyect

I assume you have installed [__Golang__](https://go.dev/dl/), [__Node.js__](https://nodejs.org/en/), [__PostgreSQL__](https://www.postgresql.org/download/), [__MongoDB__](https://www.mongodb.com/docs/manual/?_ga=2.168585697.1162239620.1651099704-1672672351.1650840520) and __Bash__. 

To test use:
```bash
bash start.sh
```

To load all services use:
```bash
bash start.sh --full
```
Both will open the documentation pages in your default browser.

##### Remember the script compiles the code "Guardian" folder should be inside $GOPATH/src!


## Documentation images

![docu](https://user-images.githubusercontent.com/104360084/165974329-2092ad39-222d-4556-8004-3d1f5e82b173.png)

![docu2](https://user-images.githubusercontent.com/104360084/165974372-5abfc5f1-d439-4e3e-b12c-aef5a4117352.png)

![docu3](https://user-images.githubusercontent.com/104360084/165974404-b7f8c2d0-f0a7-42db-b235-816b844ab2af.png)


I'm going to keep adding features to this project.

### Enjoy it!
