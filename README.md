
#### Run the Application
Here is the steps to run it with `docker-compose`


### Clone 
```bash
$ git clone https://github.com/younesi/go-url-shortener.git
```
### Move to project
```bash
$ cd go-url-shortener
```
### Build it up
```bash
docker-compose up --build -d
```
#### Check if the containers are running
```bash
$ docker ps
```

## APIs

### Postman
 - Postman file(URL shortener.postman_collection.json) is added to the repository
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
$ curl --location --request GET 'http://localhost:3000/api/v1/Rx8hdYnM'
```

#### Stop
```bash
docker-compose down
```
