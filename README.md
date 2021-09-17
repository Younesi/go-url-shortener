

#### Run the Application
Here is the steps to run it with `docker-compose`



# Clone 
$ git clone https://github.com/younesi/go-url-shortener.git

# move to project
$ cd go-url-shortener

# build it up
docker-compose up --build -d

# check if the containers are running
$ docker ps

# Welcome API
// say the curls
$ curl localhost:3000/urls

# Stop
docker-compose down

