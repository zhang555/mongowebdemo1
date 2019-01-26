# mongowebdemo1

## How to use
```
git clone git@github.com:zhang555/mongowebdemo1.git
cd mongowebdemo1
docker-compose build
docker-compose up
```

access url http://localhost/mongoweb

username | password
---      |---
admin | admin
user1 | user1
user2 | user2


## Technology stack
* docker
* docker-compose
* golang
* vue
* mongoDB
* nginx



## Feature

### Backend
* Multiple roles, different permissions for each role
* CRUD operations for MongoDB



### Frontend
* Each role sees a different sidebar
* Call the interface provided by the backend

## Picture

docker-compose build

![](https://github.com/zhang555/mongowebdemo1/blob/master/picture/docker-compose%20build.png)

docker-compose up

![](https://github.com/zhang555/mongowebdemo1/blob/master/picture/docker-compose%20up.png)

![](https://github.com/zhang555/mongowebdemo1/blob/master/picture/admin.png)

![](https://github.com/zhang555/mongowebdemo1/blob/master/picture/user1.png)



