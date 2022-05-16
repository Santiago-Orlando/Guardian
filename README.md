# Guardian

## Overview

"Guardian" is a file-saving application and is made up of 4 microservices described below.

- Authentication:
    - Login
    - Register
    - Password recovery
 
- File Storage:
    - Compresses recived files
    - Store files ( In docker File System )
    - Send files
    - Delete files
    - Prevent file duplication
    - File hashing

- Error Logger:
    - Store Errors
    - Send Email to an admin ( if any service fail )
 
- Proxy:
    - Unify all routes

#### Project description

I have used __MongoDB__ to store the users, the authentication of the users is made with a __JWT__ Token and the password recovery is achieved throw an email with a token. The email is sent using the standard library of __Golang__ "net/smtp".

The files could be sent using a multipart/form-data or a base64 and there are stored in the docker filesystem, before storing the file I load the name, hash, user_id, and the new name of the file used in the FS in a __PostgreSQL__ Database.
I Prevent the file duplication using the hash, so even if it belongs to two different users, both can access the file but it is only stored once.

The errors are stored in a __PostgreSQL__ Database and that service pings the other two services if anyone doesn't respond it sends an email to an admin.
This service is made in __JavaScript__

Proxy unifies all routes and converts a multipart/form-data file to base64. Also, have a "/docs" route where the documentation is displayed in a friendly mode.

The documentation is made with __Swagger__.

## Start the Proyect

I assume you have __docker-compose__ and __Bash__. 

To load all services use:
```bash
bash start.sh
```

## Documentation images

![docu](https://user-images.githubusercontent.com/104360084/165974329-2092ad39-222d-4556-8004-3d1f5e82b173.png)

![docu2](https://user-images.githubusercontent.com/104360084/165974372-5abfc5f1-d439-4e3e-b12c-aef5a4117352.png)

![image](https://user-images.githubusercontent.com/104360084/166463810-291e63bc-8d9e-4e7b-a4f4-c41b2fea11bd.png)


I'm going to keep adding features to this project.

### Enjoy it!
