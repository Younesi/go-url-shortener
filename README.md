
### How To Run This Project
> Make Sure you have run the urls.sql in your mysql


Since the project already use Go Module, I recommend to put the source code in any folder but GOPATH.

#### Run the Testing

```bash
$ make test
```

#### Run the Applications
Here is the steps to run it with `docker-compose`

```bash
#move to directory
$ cd workspace

# Clone into YOUR $GOPATH/src
$ git clone https://github.com/younesi/go-url-shortener.git

#move to project
$ cd go-url-shortener

# build it up
docker-compose up --build -d

# check if the containers are running
$ docker ps

# Welcome API
// say the curls
$ curl localhost:8080/urls

# Stop
docker-compose down

