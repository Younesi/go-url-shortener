

#### Run the Application
Here is the steps to run it with `docker-compose`



### Clone 
$ git clone https://github.com/younesi/go-url-shortener.git

### move to project
$ cd go-url-shortener

### build it up
docker-compose up --build -d

### check if the containers are running
$ docker ps

#### Create short URL API
```bash
$ curl --request POST \
--data '{
    "long_url": "https://www.linkedin.com/in/mahdi-younesi/"
}' \
  http://localhost:3000/api/v1/create-short-url
```
#### Get short URL API
```bash
$ curl --request GET \
  http://localhost:3000/api/v1/Rx8hdYnM
```

# Stop
docker-compose down

